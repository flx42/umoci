/*
 * umoci: Umoci Modifies Open Containers' Images
 * Copyright (C) 2016, 2017 SUSE LLC.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package main

import (
	"github.com/urfave/cli"
)

var rawSubcommand = cli.Command{
	Name:  "raw",
	Usage: "advanced internal image tooling",
	ArgsUsage: `raw <command> [<args>...]

In order to facilitate more advanced uses of umoci, the umoci-raw(1)
subcommands allow for more fine-grained information to be provided from umoci.
Please do not use these commands if you are not familiar with the intricacies
of the OCI specifications. The top-level umoci-unpack(1) and similar commands
should be sufficient for most use-cases.`,

	Subcommands: []cli.Command{
		rawConfigCommand,
		rawUnpackCommand,
	},
}
