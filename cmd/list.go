// Copyright © 2016 davey mcglade  <davey.mcglade@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"

	"github.com/daveym/lint/lintapi"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List items stored witin your Pocket data store",
	Long: `List provides you with the means to query Pocket, pull back items by specific order 
	and then use the results to perform follow up actions such as deleting or tagging`,
	Run: func(cmd *cobra.Command, args []string) {

		var ItemRequest LintApi.ItemRequest
		ItemRequest.ConsumerKey = ConsumerKey
		ItemRequest.AccessToken = AccessToken
		ItemRequest.State = LintApi.ItemStateAll
		ItemRequest.ContentType = string(LintApi.ContentTypeArticle)
		ItemRequest.DetailType = string(LintApi.DetailTypeSimple)
		ItemRequest.Count = 10

		items, err := LintApi.GetItems(ItemRequest)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(items.GivenTitle)
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
