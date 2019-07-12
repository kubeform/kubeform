package main

//func generateCRDDefinitions() {
//	apis.EnableStatusSubresource = true
//
//	filename := gort.GOPath() + "/src/kubeform.dev/kubeform/apis/aws/v1alpha1/crds.yaml"
//	os.Remove(filename)
//
//	err := os.MkdirAll(filepath.Join(gort.GOPath(), "/src/kubeform.dev/kubeform/api/crds"), 0755)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	crds := []*crd_api.CustomResourceDefinition{
//		catalogv1alpha1.PostgresVersion{}.CustomResourceDefinition(),
//	}
//	for _, crd := range crds {
//		filename := filepath.Join(gort.GOPath(), "/src/kubeform.dev/kubeform/api/crds", crd.Spec.Names.Singular+".yaml")
//		f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
//		if err != nil {
//			log.Fatal(err)
//		}
//		crdutils.MarshallCrd(f, crd, "yaml")
//		f.Close()
//	}
//}
//
//func generateSwaggerJson() {
//	var (
//		Scheme = runtime.NewScheme()
//		Codecs = serializer.NewCodecFactory(Scheme)
//	)
//
//	kubedbinstall.Install(Scheme)
//	cataloginstall.Install(Scheme)
//	authorizationinstall.Install(Scheme)
//
//	apispec, err := openapi.RenderOpenAPISpec(openapi.Config{
//		Scheme: Scheme,
//		Codecs: Codecs,
//		Info: spec.InfoProps{
//			Title:   "KubeDB",
//			Version: "v0",
//			Contact: &spec.ContactInfo{
//				Name:  "AppsCode Inc.",
//				URL:   "https://appscode.com",
//				Email: "hello@appscode.com",
//			},
//			License: &spec.License{
//				Name: "Apache 2.0",
//				URL:  "https://www.apache.org/licenses/LICENSE-2.0.html",
//			},
//		},
//		OpenAPIDefinitions: []common.GetOpenAPIDefinitions{
//			catalogv1alpha1.GetOpenAPIDefinitions,
//		},
//		Resources: []openapi.TypeInfo{
//			{catalogv1alpha1.SchemeGroupVersion, catalogv1alpha1.ResourcePluralPostgresVersion, catalogv1alpha1.ResourceKindPostgresVersion, false},
//		},
//	})
//	if err != nil {
//		glog.Fatal(err)
//	}
//
//	filename := gort.GOPath() + "/src/kubeform.dev/kubeform/api/openapi-spec/swagger.json"
//	err = os.MkdirAll(filepath.Dir(filename), 0755)
//	if err != nil {
//		glog.Fatal(err)
//	}
//	err = ioutil.WriteFile(filename, []byte(apispec), 0644)
//	if err != nil {
//		glog.Fatal(err)
//	}
//}

func main() {
	//generateCRDDefinitions()
	//generateSwaggerJson()
}
