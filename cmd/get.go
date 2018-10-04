package cmd

import (
	"encoding/json"
	"fmt"

	//"github.com/ghodss/yaml"
	"github.com/go-resty/resty"
	//"github.com/olekukonko/tablewriter"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/containership/csctl/pkg/cloud/api/types"
)

// Flags
var (
	outputFormat string
)

func getResource(url string) (string, error) {
	token := viper.GetString("token")
	if token == "" {
		return "", errors.New("please provide a token")
	}

	authHeader := fmt.Sprintf("JWT %s", token)

	resp, err := resty.R().SetHeader("Authorization", authHeader).Get(url)
	if err != nil {
		return "", errors.Wrap(err, "error requesting resource")
	}

	return string(resp.Body()), err
}

func getOrganizations() (*[]types.Organization, error) {
	url := fmt.Sprintf("%s/v3/organizations", viper.Get("apiBaseURL"))

	token := viper.GetString("token")
	if token == "" {
		return nil, errors.New("please provide a token")
	}

	authHeader := fmt.Sprintf("JWT %s", token)

	orgs := make([]types.Organization, 0)
	resp, err := resty.R().SetHeader("Authorization", authHeader).
		SetResult(orgs).
		Get(url)

	if err != nil {
		return nil, errors.Wrap(err, "error requesting resource")
	}

	if resp.IsError() {
		return nil, errors.Errorf("Containership Cloud responded with status %d: %s", resp.StatusCode(), resp.Body())
	}

	return resp.Result().(*[]types.Organization), err
}

func getClusters(organizationID string) (string, error) {
	url := fmt.Sprintf("%s/v3/organizations/%s/clusters", viper.Get("apiBaseURL"), organizationID)
	return getResource(url)
}

func getNodePools(organizationID, clusterID string) (string, error) {
	url := fmt.Sprintf("%s/v3/organizations/%s/clusters/%s/node-pools",
		viper.Get("provisionBaseURL"), organizationID, clusterID)
	return getResource(url)
}

// TODO this function is beyond terrible
func outputResponse(resp interface{}) {
	switch outputFormat {
	case "", "table":
		// Default
		// TODO just dump raw response (json blob) for now
		fmt.Println(resp)

	case "json":
		j, err := json.MarshalIndent(resp, "", "  ")
		if err != nil {
			fmt.Printf("Error formatting JSON: %v", err)
			return
		}

		fmt.Println(string(j))

	case "yaml":
		//y, err := yaml.JSONToYAML([]byte(resp))
		//if err != nil {
		//fmt.Println("(Error converting to YAML)")
		//return
		//}

		//fmt.Println(string(y))

	case "jsonpath":
		fallthrough
	default:
		// TODO
		fmt.Printf("(output format %s not supported)", outputFormat)
	}
}

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get <resource>",
	Short: "Get a resource",
	Long: `Get a resource

TODO this is a long description`,

	Args: cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		resource := args[0]
		switch resource {
		case "org", "orgs", "organization", "organizations":
			resp, err := getOrganizations()
			if err != nil {
				fmt.Println(err)
				return
			}

			outputResponse(resp)

		case "cluster", "clusters":
			if organizationID == "" {
				fmt.Println("organization is required")
				return
			}

			resp, err := getClusters(organizationID)
			if err != nil {
				fmt.Println(err)
				return
			}

			outputResponse(resp)

		case "nodepool", "nodepools", "np", "nps":
			if organizationID == "" || clusterID == "" {
				fmt.Println("organization and cluster are required")
				return
			}

			resp, err := getNodePools(organizationID, clusterID)
			if err != nil {
				fmt.Println(err)
				return
			}

			outputResponse(resp)

		default:
			fmt.Println("Error: invalid resource specified: %q", resource)
		}
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	getCmd.Flags().StringVarP(&outputFormat, "output", "o", "", "Output format")
}
