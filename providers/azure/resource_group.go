// Copyright 2019 The Terraformer Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package azure

import (
	"context"
	"log"

	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2019-05-01/resources"
	"github.com/Azure/go-autorest/autorest/azure/auth"
	"github.com/GoogleCloudPlatform/terraformer/terraform_utils"
)

type ResourceGroupGenerator struct {
	AzureService
}

func (g ResourceGroupGenerator) createResources(groupListResultIterator resources.GroupListResultIterator) []terraform_utils.Resource {
	var resources []terraform_utils.Resource
	for groupListResultIterator.NotDone() {
		group := groupListResultIterator.Value()
		resources = append(resources, terraform_utils.NewSimpleResource(
			*group.ID,
			*group.Name,
			"azurerm_resource_group",
			"azurerm",
			[]string{}))
		err := groupListResultIterator.Next()
		if err != nil {
			log.Println(err)
			break
		}
	}
	return resources
}

func (g *ResourceGroupGenerator) InitResources() error {
	ctx := context.Background()
	groupsClient := resources.NewGroupsClient(g.Args["subscription"].(string))
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		return err
	}
	groupsClient.Authorizer = authorizer
	output, err := groupsClient.ListComplete(ctx, "", nil)
	if err != nil {
		return err
	}
	g.Resources = g.createResources(output)
	return nil
}
