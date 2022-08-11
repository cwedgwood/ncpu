
test: .iid
	docker run --rm $$(cat .iid)

n: Makefile n.go
	CGOENABLED=0 go build -ldflags="-w -s" -o $@

.iid: Dockerfile Makefile n
	docker build --iidfile=$@ .


clean:
	rm -f *~ n .iid
