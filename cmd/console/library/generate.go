package library

//go:generate go-bindata -nometadata -pkg library -o bindata.go bignumber.js neb-light.js
//go:generate goimports -w bindata.go
