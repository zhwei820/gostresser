install_dep:
	go get .; cd ui; yarn install; cd -;

build_web:
	go build .

build_worker:
	cd worker; go build .; cd -

build_ui:
	cd ui; yarn install; yarn build; cd -

build_image:
	docker build . -t daocloud.io/zhwei820/gostresser

build_go: build_worker build_web 

serve_api:
	./gostresser

serve_ui:
	cd ui; yarn serve

update_image: build_go build_ui build_image
