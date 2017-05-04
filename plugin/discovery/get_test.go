package discovery

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVersionListing(t *testing.T) {
	handler := http.NewServeMux()
	handler.HandleFunc("/terraform-providers/terraform-provider-test/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(versionList))
	})

	server := httptest.NewServer(handler)
	defer server.Close()

	providersURL.releases = server.URL + "/"

	versions, err := listProviderVersions("test")
	if err != nil {
		t.Fatal(err)
	}

	expectedSet := map[string]bool{
		"1.2.4": true,
		"1.2.3": true,
		"1.2.1": true,
	}

	for _, v := range versions {
		if !expectedSet[v.String()] {
			t.Fatalf("didn't get version %s in listing")
		}
		delete(expectedSet, v.String())
	}
}

const versionList = `<!DOCTYPE html>
<html>
<body>
  <ul>
  <li>
    <a href="../">../</a>
  </li>
  <li>
    <a href="/terraform-provider-test/1.2.3/">terraform-provider-test_1.2.3</a>
  </li>
  <li>
    <a href="/terraform-provider-test/1.2.1/">terraform-provider-test_1.2.1</a>
  </li>
  <li>
    <a href="/terraform-provider-test/1.2.4/">terraform-provider-test_1.2.4</a>
  </li>
  </ul>
  <footer>
    Proudly fronted by <a href="https://fastly.com/?utm_source=hashicorp" target="_TOP">Fastly</a>
  </footer>
</body>
</html>
`
