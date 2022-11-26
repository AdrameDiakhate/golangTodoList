def 

test_api:

	echo "Testing the api..."


build_api:
	
	docker login -u="adrame98" -p="adrame@2022"

	docker tag SOURCE_IMAGE[:TAG] TARGET_IMAGE[:TAG]

	docker build [OPTIONS] PATH | URL | -

	docker push adrame98/golangtodolist:tagname
PushDev
deploy_api: