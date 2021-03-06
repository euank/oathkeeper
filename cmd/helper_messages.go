// Copyright © 2017-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package cmd

import (
	"fmt"
	"os"
)

var corsMessage = `CORS CONTROLS
==============
- CORS_ALLOWED_ORIGINS: A list of origins (comma separated values) a cross-domain request can be executed from.
	If the special * value is present in the list, all origins will be allowed. An origin may contain a wildcard (*)
	to replace 0 or more characters (i.e.: http://*.domain.com). Usage of wildcards implies a small performance penality.
	Only one wildcard can be used per origin. The default value is *.
	Example: CORS_ALLOWED_ORIGINS=http://*.domain.com,http://*.domain2.com

- CORS_ALLOWED_METHODS: A list of methods  (comma separated values) the client is allowed to use with cross-domain
	requests. Default value is simple methods (GET and POST).
	Example: CORS_ALLOWED_METHODS=POST,GET,PUT

- CORS_ALLOWED_CREDENTIALS: Indicates whether the request can include user credentials like cookies, HTTP authentication
	or client side SSL certificates. The default is false.

- CORS_DEBUG: Debugging flag adds additional output to debug server side CORS issues.

- CORS_MAX_AGE: Indicates how long (in seconds) the results of a preflight request can be cached. The default is 0 which stands for no max age.

- CORS_ALLOWED_HEADERS: A list of non simple headers (comma separated values) the client is allowed to use with cross-domain requests.

- CORS_EXPOSED_HEADERS: Indicates which headers (comma separated values) are safe to expose to the API of a CORS API specification.`

var databaseUrl = `- DATABASE_URL: A URL to a persistent backend. Hydra supports various backends:
  - Memory: If DATABASE_URL is "memory", data will be written to memory and is lost when you restart this instance.
  	Example: DATABASE_URL=memory

  - Postgres: If DATABASE_URL is a DSN starting with postgres:// PostgreSQL will be used as storage backend.
	Example: DATABASE_URL=postgres://user:password@host:123/database

	If PostgreSQL is not serving TLS, append ?sslmode=disable to the url:
	DATABASE_URL=postgres://user:password@host:123/database?sslmode=disable

  - MySQL: If DATABASE_URL is a DSN starting with mysql:// MySQL will be used as storage backend.
	Example: DATABASE_URL=mysql://user:password@tcp(host:123)/database?parseTime=true

	Be aware that the ?parseTime=true parameter is mandatory, or timestamps will not work.`

func fatalf(msg string, args ...interface{}) {
	fmt.Printf(msg+"\n", args...)
	os.Exit(1)
}
