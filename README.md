# Go Talks

[![Go Reference](https://pkg.go.dev/badge/golang.org/x/talks.svg)](https://pkg.go.dev/golang.org/x/talks)

This repository holds various Go talks that may be viewed with the present tool.

## Viewing Locally

To install the present tool, use `go get`:

```
go get golang.org/x/tools/cmd/present
```

To view talks, run `present` in the `content` directory and then
visit [http://127.0.0.1:3999](http://127.0.0.1:3999) in your browser.

## Report Issues / Send Patches

This repository uses Gerrit for code changes. To learn how to submit changes to
this repository, see https://golang.org/doc/contribute.html.

The main issue tracker for the talks repository is located at
https://github.com/golang/go/issues. Prefix your issue with "x/talks:" in the
subject line, so it is easy to find.

## Deploying

To deploy these talks to talks.golang.org hosted on App Engine:

1.	Add a version of `golang.org/x/tools/cmd/present` that you intend to
	use in this deploy as a dependency of this module:

	```
	go get -d golang.org/x/tools/cmd/present@latest
	```

2.	Copy the contents of `golang.org/x/tools/cmd/present` to the root of
	this repository:

	```
	rsync -r $(go list -f '{{.Dir}}' golang.org/x/tools/cmd/present)/ .
	```

3.	Run the deploy:

	```
	GO111MODULE=on gcloud --project=golang-org app deploy --no-promote app.yaml
	```

	This will create a new version, which can be viewed within the
	[golang-org GCP project](https://console.cloud.google.com/appengine/versions?project=golang-org&serviceId=talks).

4.	Check that the deployed version looks OK (click the version link in GCP).

5.	If all is well, click "Migrate Traffic" to move 100% of the talks.golang.org
	traffic to the new version.

6.	Clean up the local file modifications that steps 1 and 2 caused.

7.	You're done.
