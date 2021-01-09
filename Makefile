generate:
	go generate

bundle_templates:
	bindata -pkg templates -o templates/template_data.go ./templates/project/...
