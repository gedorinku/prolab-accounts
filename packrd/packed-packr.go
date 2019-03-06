// +build !skippackr
// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "5084aee7b21f3b7cdb10b195a9522dc4"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"7ceeb028aedf0c11d541e1e265dc4a09": "1f8b08000000000000ff948f414ac3501086f7ef145ea0edde337415e801a4142dd256dad75529383356840604a1b4c542c545883cd468408c09b9cc9f18730bc92317c8ac66f1cd377c902de415f20239809fc106ec4142c81dc8fcae1ff37405da805d500ada828eb866a5b2eff59fbfeb395d90b177efe0181c59c483789004fc05498a7d5cba9fe087fccdcd7e6e6b80ee414f95986eac6fb1683b83f3e14c4fcff47032ee39dde552a956b35159b42a36fbe2232ccd2e0ff8a4aee3c0061eabae7a3f94e2ab0badaf66a79dcee57c3a1f0d5ae3be6ef727a3c65fff030000ffffacd310e344010000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("emails", "./emails")
		b.SetResolver("invitation.tmpl", packr.Pointer{ForwardBox: gk, ForwardPath: "7ceeb028aedf0c11d541e1e265dc4a09"})
	}()

	return nil
}()
