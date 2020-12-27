package miru_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/influx6/npkg/nerror"
	"github.com/stretchr/testify/require"

	"github.com/influx6/groundlayer/pkg/miru"
	"github.com/influx6/groundlayer/pkg/miru/parse"
)

var (
	allFuncs = []map[string]string{
		{
			"rewrap": "rewrap",
		},
	}

	partialFiles = &memoryFS{
		files: map[string]string{
			"some.tmper": `
				{{define "some.tmper" }}
					This is some.tmper text
				{{end}}

				{{define "some" }}
					This is some text
				{{end}}

				{{define "some2" string }}
					This is some2 text
				{{end}}
			`,
			"thing.tmper": `
				Hello, welcome
				{{template "some.tmper#some"}}
				{{template "some.tmper"}}
			`,
		},
	}
)

func TestTextWriter_FileSystem(t *testing.T) {
	var targetFile, err = partialFiles.GetFile("thing.tmper")
	require.NoError(t, err)
	require.NotNil(t, targetFile)

	var parsed, parsedErr = targetFile.Parse(miru.DefaultOption, nil)
	require.NoError(t, parsedErr)
	require.NotNil(t, parsed)

	var content, contentErr = parsed.Format("test_pack")
	require.NoError(t, contentErr)

	hasNotText(t, content, `func definedTemplate2(ctx some.tmper, rootDoc *domu.Node)`)
	hasText(t, content, `func definedTemplate2(page *peji.Page, ctx interface{}, rootDoc *domu.Node)`)
	hasText(t, content, `This is some text`)
	hasText(t, content, `definedTemplate2(ctx, rootDoc)`)
	hasText(t, content, `func definedTemplate5(page *peji.Page, ctx interface{}, rootDoc *domu.Node)`)
	hasText(t, content, `This is some.tmper text`)
	hasText(t, content, `definedTemplate5(ctx, rootDoc)`)
}

func TestTextWriter_ParseTextTemplateWithMethodsWithUsage(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{model Address | Zip:String, Street: string, Number:Int }}
		{{model Product | Isin:String, Price: Float, Address:Address }}
		{{modelType ProductMapping |  Map(string, Product) }}
		
		{{method renderSubProducts | product:Product, subProducts:ProductMapping }}
		  <div class="product">
			{{ product.PageName | print }}
			{{if len subProducts}}
			  <div class="item">
				<span class="price">No sub products available</span>
			  </div>
			{{else}}
			  <div class="item">
				<span class="price">Categories</span>
			  </div>
			{{end}}
		  </div>
		{{endMethod}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `type Address struct`)
	hasText(t, builder, `type Product struct`)
	hasText(t, builder, `type ProductMapping map[string]Product`)
	hasText(t, builder, `func renderSubProducts`)
	hasText(t, builder, `func RenderLayout(page *peji.Page, ctx string, rootDoc *domu.Node)`)
	hasText(t, builder, `helpers.Print((product).PageName)`)
}

func TestTextWriter_ParseTextTemplateWithMethods(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{model Address | Zip:String, Street: string, Number:Int }}
		{{model Product | Isin:String, Price: Float, Address:Address }}
		{{modelType ProductMapping |  Map(string, Product) }}
		
		{{method renderSubProducts | product:Product, subProducts:ProductMapping }}
		  <div class="product">
			{{if len subProducts}}
			  <div class="item">
				<span class="price">No sub products available</span>
			  </div>
			{{else}}
			  <div class="item">
				<span class="price">Categories</span>
			  </div>
			{{end}}
		  </div>
		{{endMethod}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `type Address struct`)
	hasText(t, builder, `type Product struct`)
	hasText(t, builder, `type ProductMapping map[string]Product`)
	hasText(t, builder, `func renderSubProducts`)
}

func TestTextWriter_ParseTextTemplateWithMethodsUsingKomponent(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{model Address | Zip:String, Street: string, Number:Int }}
		{{model Product | Isin:String, Price: Float, Address:Address }}
		{{modelType ProductMapping |  Map(string, Product) }}

		{{method renderSubProducts | product:Product, subProducts:ProductMapping }}
		  <div class="product">
			{{if len subProducts}}
			  <div class="item">
				<span class="price">No sub products available</span>
			  </div>
			{{else}}
			  <div class="item">
				<span class="price">Categories</span>
			  </div>
			{{end}}
		  </div>
		{{endMethod}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `type Address struct`)
	hasText(t, builder, `type Product struct`)
	hasText(t, builder, `type ProductMapping map[string]Product`)
	hasText(t, builder, `func renderSubProducts`)
	hasText(t, builder, `func RenderLayout(page *peji.Page, ctx string, rootDoc *domu.Node)`)
}

func TestTextWriter_ParseTextTemplateWithMethodsUsingKomponentAndMountsFailed(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{model Address | Zip:String, Street: string, Number:Int }}
		{{model Product | Isin:String, Price: Float, Address:Address }}
		{{modelType ProductMapping |  Map(string, Product) }}

		{{method renderSubProducts | product:Product, subProducts:ProductMapping }}
			<div class="product">
				<span class="price">No sub products available</span>
			</div>
			{{ mount UserCars }}
		{{endMethod}}
	`, true)

	require.Empty(t, builder)
}

func TestTextWriter_ParseTextTemplateWithSubDOmMount(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{model Address | Zip:String, Street: string, Number:Int }}
		{{model Product | Isin:String, Price: Float, Address:Address }}
		{{modelType ProductMapping |  Map(string, Product) }}

		{{method renderSubProducts | product:Product, subProducts:ProductMapping }}
			<div class="product">
				<span class="price">No sub products available</span>
				{{ mount UserCars product }}
			</div>
		{{endMethod}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `type Address struct`)
	hasText(t, builder, `type Product struct`)
	hasText(t, builder, `type ProductMapping map[string]Product`)
	hasText(t, builder, `func renderSubProducts`)
	hasText(t, builder, `if page.HasStatic("UserCars") {`)
	hasText(t, builder, `page.Static("UserCars").Render(product).Mount(node6)`)
	hasText(t, builder, `renderSubProducts(page *peji.Page, product Product, subProducts ProductMapping, parentNode *domu.Node) {`)
}

func TestTextWriter_ParseTextTemplateWithMethodsUsingKomponentAndMounts(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{model Address | Zip:String, Street: string, Number:Int }}
		{{model Product | Isin:String, Price: Float, Address:Address }}
		{{modelType ProductMapping |  Map(string, Product) }}

		{{method renderSubProducts | product:Product, subProducts:ProductMapping }}
			<div class="product">
				<span class="price">No sub products available</span>
			</div>
			{{ mount UserCars product }}
		{{endMethod}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `type Address struct`)
	hasText(t, builder, `type Product struct`)
	hasText(t, builder, `type ProductMapping map[string]Product`)
	hasText(t, builder, `func renderSubProducts`)
	hasText(t, builder, `if page.HasStatic("UserCars") {`)
	hasText(t, builder, `page.Static("UserCars").Render(product).Mount(parentNode)`)
	hasText(t, builder, `renderSubProducts(page *peji.Page, product Product, subProducts ProductMapping, parentNode *domu.Node) {`)
}

func TestTextWriter_ParseTextWithModelAndModelType(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{model Address | Zip:string, Street: string, Number:Int }}
		{{model Product | isin:String, price: Float, address:Address }}
		{{modelType ProductLog | Any }}
		{{modelType ProductLog2 | any }}
		{{modelType ProductAddressCode | String }}
		{{modelType ProductMapping |  Map(string, Product) }}
		{{modelType ListProductMapping |   List(ProductMapping) }}
		{{modelType ProductList | List(Product) }}
		{{modelType ProductMapping2 |  Map(String, Product) }}
		{{modelType ProductMapping3 |  Map(String, Any) }}
		{{model Store | Address:string, Date:Time, Products:ProductList, Mappings:ProductMapping }}

		{{rootType Store }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `type Address struct`)
	hasText(t, builder, `type Product struct`)
	hasText(t, builder, `type Store struct`)
	hasText(t, builder, `type ProductAddressCode string`)
	hasText(t, builder, `type ListProductMapping []ProductMapping`)
	hasText(t, builder, `type ProductLog interface{}`)
	hasText(t, builder, `type ProductLog2 interface{}`)
	hasText(t, builder, `type ProductMapping map[string]Product`)
	hasText(t, builder, `type ProductAddressCode string`)
	hasText(t, builder, `type ProductMapping3 map[string]interface{}`)
}

func TestTextWriter_ParseMountWithRootType(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ rootType string }}
		{{ mount user }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, ctx string, rootDoc *domu.Node) {`)
	hasText(t, builder, `if page.HasStatic("user") {`)
	hasText(t, builder, `page.Static("user").Render(ctx).Mount(rootDoc)`)
}

func TestTextWriter_ParseMountListWithRootType(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ rootType string }}
		{{ mountList user }}
	`, false)

	fmt.Println(builder)
	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, ctx string, rootDoc *domu.Node) {`)
	hasText(t, builder, `if page.HasStatic("user") {`)
	hasText(t, builder, `page.Static("user").Render(ctx).Mount(rootDoc)`)
}

func TestTextWriter_ParseMount(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ mount user }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `if page.HasStatic("user") {`)
	hasText(t, builder, `page.Static("user").Render(ctx).Mount(rootDoc)`)
}

func TestTextWriter_ParseMountLiveList(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ mountLiveList /profile | profile . }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `if page.HasLive("profile") {`)
	hasText(t, builder, `page.Live("/profile", "profile").Render(ctx).Mount(rootDoc)`)
}

func TestTextWriter_ParseMountLive(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ mountLive /profile | profile . }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `if page.HasLive("profile") {`)
	hasText(t, builder, `page.Live("/profile", "profile").Render(ctx).Mount(rootDoc)`)
}

func TestTextWriter_ParseMountWithFields(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ $x := .Field }}
		
		{{ mount user .Field1.Field2 }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `if page.HasStatic("user") {`)
	hasText(t, builder, `page.Static("user").Render(ctx.Field1.Field2).Mount(rootDoc)`)
}

func TestTextWriter_ParseMountWithArgumentVariable(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ $x := .Field }}
		
		{{ mount user $x }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `if page.HasStatic("user") {`)
	hasText(t, builder, `page.Static("user").Render(x).Mount(rootDoc)`)
}

func TestTextWriter_ParseMountWithString(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ $x := .Field }}
		
		{{ mount user  "alex" }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `if page.HasStatic("user") {`)
	hasText(t, builder, `page.Static("user").Render("alex").Mount(rootDoc)`)
}

func TestTextWriter_ParseTextString(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{define "footer" string }}
			This is rendered text
		{{end}}

		{{rootType string }}
		{{template "footer" .}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, ctx string, rootDoc *domu.Node) {`)
	hasText(t, builder, `func definedTemplate3(page *peji.Page, ctx string, rootDoc *domu.Node) {`)
}

func TestTextWriter_ParseTextTemplate(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{modelType ItemList | List(string) }}
		{{model Product | Items: ItemList }}
		{{rootType Product }}

		{{define "footer.tmper" Product }}
			This is rendered text
			{{range .Items}}
				<div class="item">{{.}}</div>
			{{end}}
		{{end}}

		{{template "footer.tmper"}}
		{{template "footer.tmper" .}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `type ItemList []string`)
	hasText(t, builder, `Items ItemList`)
	hasText(t, builder, `func RenderLayout(page *peji.Page, ctx Product, rootDoc *domu.Node) {`)
	hasText(t, builder, `func definedTemplate5(page *peji.Page, ctx Product, rootDoc *domu.Node) {`)
}

func TestTextWriter_ParseDocType(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		<!DocType  html PUBLIC "-//W3C//DTD XHTML 1.1//EN" "http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd">
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `html PUBLIC \"-//W3C//DTD XHTML 1.1//EN\" \"http://www.w3.org/TR/xhtml11/DTD/xhtml11.dtd\""`)
}

func TestTextWriter_ParseLayoutArgument(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}
		<p>It was you</p>
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, ctx string, rootDoc *domu.Node) {`)
}

func TestTextWriter_ParseLayoutArgumentWithNoRoot(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		<p>It was you</p>
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
}

func TestTextWriter_ParseWithPackageImport(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		<p>It was you</p>
		{{ import plackers | github.com/influx6/groundlayer/pkg/plackers }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
}

func TestTextWriter_ParseWithPackageUse(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ #helpers.Printf . }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
}

func TestTextWriter_ParseWithPackageUseWithField(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ @helpers.TotalLives }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.AttachToNode(helpers.TotalLives, rootDoc)`)
	hasText(t, builder, `helpers.TotalLives`)
}

func TestTextWriter_ParseWithPackageUseWithFieldWithPipe(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ @helpers.TotalLives | print }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.Print(helpers.TotalLives)`)
	hasText(t, builder, `helpers.AttachToNode(helpers.Print(helpers.TotalLives), rootDoc)`)
}

func TestTextWriter_ParseFieldWithImportedFieldWithPipe(t *testing.T) {
	var builder = textWriterTestHelper(t, `{{ @helpers.Count | print }}`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.Print(helpers.Count)`)
	hasText(t, builder, `helpers.AttachToNode(helpers.Print(helpers.Count), rootDoc)`)
}

func TestTextWriter_ParseFieldWithImportedFieldWithPipe2(t *testing.T) {
	var builder = textWriterTestHelper(t, `{{ @helpers.Count | print .Field1 }}`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.Print(helpers.Count, ctx.Field1)`)
	hasText(t, builder, `helpers.AttachToNode(helpers.Print(helpers.Count, ctx.Field1), rootDoc)`)
}

func TestTextWriter_ParseUsingPrint(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ "alex" | print dom | noop }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.Print("alex", rootDoc)`)
}

func TestTextWriter_ParseUsingPrint2(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ print "alex" dom }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.Print("alex", rootDoc)`)
	hasText(t, builder, `helpers.AttachToNode(helpers.Print("alex", rootDoc), rootDoc)`)
}

func TestTextWriter_ParseUsingDOMDirective(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ "alex" | #helpers.AddToDOM dom }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.AddToDOM("alex", rootDoc)`)
	hasText(t, builder, `helpers.AttachToNode(helpers.AddToDOM("alex", rootDoc), rootDoc)`)
}

func TestTextWriter_ParseUsingDOMDirective2(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ #helpers.AddToDOM "alex" dom }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.AddToDOM("alex", rootDoc)`)
	hasText(t, builder, `helpers.AttachToNode(helpers.AddToDOM("alex", rootDoc), rootDoc)`)
}

func TestTextWriter_ParseUsingDOMDirective3(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ "alex" | #helpers.AddToDOM dom | noop }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.AddToDOM("alex", rootDoc)`)
	hasNotText(t, builder, `helpers.AttachToNode(helpers.AddToDOM("alex", rootDoc), rootDoc)`)
}

func TestTextWriter_ParseUsingDOMDirective4(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ #helpers.AddToDOM "alex" dom | noop }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.AddToDOM("alex", rootDoc)`)
	hasNotText(t, builder, `helpers.AttachToNode(helpers.AddToDOM("alex", rootDoc), rootDoc)`)
}

func TestTextWriter_ParseUsingDOMDirective5(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		<div>
			<p>
				{{ #helpers.AddToDOM "alex" dom | print | noop }}
			</p>
		</div>
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.Print(helpers.AddToDOM("alex", node4))`)
	hasNotText(t, builder, `helpers.AttachToNode(helpers.Print(helpers.AddToDOM("alex", rootDoc)), rootDoc)`)
}

func TestTextWriter_ParseUsingDOMDirective6(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		<div>
			<p>
				{{ #helpers.AddToDOM "alex" dom | print }}
			</p>
		</div>
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.Print(helpers.AddToDOM("alex", node4))`)
	hasText(t, builder, `helpers.AttachToNode(helpers.Print(helpers.AddToDOM("alex", node4)), node4)`)
}

func TestTextWriter_ParseUsingDOMDirective7(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		<div>
			<p>
				{{ printf "alex" dom | print }}
			</p>
		</div>
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.Print(helpers.Printf("alex", node4))`)
	hasText(t, builder, `helpers.AttachToNode(helpers.Print(helpers.Printf("alex", node4)), node4)`)
}

func TestTextWriter_ParseWithPackageUseAsFieldTwoVar(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ print @helpers.TotalLives "alex" }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.Print(helpers.TotalLives, "alex")`)
	hasText(t, builder, `helpers.AttachToNode(helpers.Print(helpers.TotalLives, "alex"), rootDoc)`)
}

func TestTextWriter_ParseWithPackageUseAsFieldTwoVar2(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ print @helpers.TotalLives "alex" | print }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.Print(helpers.TotalLives, "alex")`)
	hasText(t, builder, `helpers.AttachToNode(helpers.Print(helpers.TotalLives, "alex"), rootDoc)`)
}

func TestTextWriter_ParseWithPackageUseAsFieldTwoVar3(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ $x := "wonder" }}
		{{ print @helpers.TotalLives "alex" $x | print }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.Print(helpers.TotalLives, "alex", x)`)
	hasText(t, builder, `helpers.AttachToNode(helpers.Print(helpers.TotalLives, "alex", x), rootDoc)`)
}

func TestTextWriter_ParseWithPackageUseAsFieldTwoVar4(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ $x := "wonder" }}
		{{ $x = (print @helpers.TotalLives "alex" $x) | print }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.Print(helpers.TotalLives, "alex", x)`)
	hasText(t, builder, `x = helpers.Print(helpers.Print(helpers.TotalLives, "alex", x))`)
}

func TestTextWriter_ParseWithPackageUseAsFieldTwoVar5(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ $x := "wonder" }}
		{{ $x = (print @helpers.TotalLives "alex" $x | print) }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.Print(helpers.TotalLives, "alex", x)`)
	hasText(t, builder, `x = helpers.Print(helpers.TotalLives, "alex", x)`)
}

func TestTextWriter_ParseWithPackageUseAsFieldTwoVar6(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ $x := "wonder" }}
		{{ $x = ((print @helpers.TotalLives "alex" $x) | print) }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.Print(helpers.TotalLives, "alex", x)`)
	hasText(t, builder, `x = helpers.Print(helpers.Print(helpers.TotalLives, "alex", x))`)
}

func TestTextWriter_ParseWithPackageUsePipe(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ .Field1 | #helpers.Printf | cap }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.Cap(helpers.Printf(ctx.Field1))`)
}

func TestTextWriter_ParseWithPackageUsePipe2(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ @helper.Field1 | #helpers.Printf | cap }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.Cap(helpers.Printf(helper.Field1))`)
}

func TestTextWriter_ParseWithPackageUsePipe3(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ .Field1 | #helpers.Printf | cap | print | #helpers.Printf }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.Printf(helpers.Print(helpers.Cap(helpers.Printf(ctx.Field1))))`)
}

func TestTextWriter_ParseWithPackageUseAsFieldInPipe(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ .Field1 | #helpers.Printf }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
}

func TestTextWriter_ParseWithPackageUseWithVariablePipe(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ $date := ( .Field1 | #slacker.Printf  ) }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
}

func TestTextWriter_ParseWithPackageUseWithVariable(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ $date := ( #slucker.Printf .Field1 .Field2 ) }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `date := slucker.Printf(ctx.Field1, ctx.Field2)`)
}

func TestTextWriter_MultiArgumentsWithImported(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ #slacker.Printf .Field1 .Field2 }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.AttachToNode(slacker.Printf(ctx.Field1, ctx.Field2), rootDoc)`)
}

func TestTextWriter_MultiArgumentsWithImportedWithPipe(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ $x := "alex" }}
		{{ $x | #slacker.Printf .Field1 .Field2 }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.AttachToNode(slacker.Printf(x, ctx.Field1, ctx.Field2), rootDoc)`)
}

func TestTextWriter_ArgPipe(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ #slacker.Printf .Field1 | print }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.AttachToNode(helpers.Print(slacker.Printf(ctx.Field1)), rootDoc`)
}

func TestTextWriter_TwoArgsPipe(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{  .Field3 .Field1 | #slacker.Printf | print }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.AttachToNode(helpers.Print(slacker.Printf(ctx.Field3, ctx.Field1)), rootDoc)`)
}

func TestTextWriter_MultipleArgumentsLaced(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{  .Field1 | #slacker.Printf .Field2 .Field3 | print }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.AttachToNode(helpers.Print(slacker.Printf(ctx.Field1, ctx.Field2, ctx.Field3)), rootDoc`)
}

func TestTextWriter_MultipleArgumentsLacedWithVar(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ $x := "alex" }}
		{{  $x .Field1 | #slacker.Printf .Field2 .Field3 | print }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.AttachToNode(helpers.Print(slacker.Printf(x, ctx.Field1, ctx.Field2, ctx.Field3)), rootDoc`)
}

func TestTextWriter_BadPipeUsage5_Error(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{  .Field3 | .Field1 | #slacker.Printf | print }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `func RenderLayout(page *peji.Page, rootDoc *domu.Node) {`)
	hasText(t, builder, `helpers.AttachToNode(helpers.Print(slacker.Printf(ctx.Field3, ctx.Field1)), rootDoc)`)
}

func TestTextWriter_ParseTextWithAndPipe(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{with .Field1 | cap "boring"}}
			{{printf "%q" . | printf "%s"}}
		{{else}}
			{{printf "%d %d %d" 11 11 11}}
		{{end}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `if helpers.IsInvalid(item3) {`)
	hasText(t, builder, `item3 := helpers.Cap(ctx.Field1, "boring")`)
	hasText(t, builder, `helpers.AttachToNode(helpers.Printf("%d %d %d", 11, 11, 11), rootDoc)`)
	hasText(t, builder, `helpers.AttachToNode(helpers.Printf(helpers.Printf("%q", item3), "%s"), rootDoc)`)
}

func TestTextWriter_ParseTextUsingWith3(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{with .Field | cap "boring" | cap "water" }}
			{{printf "%q" . | printf "%s"}}
		{{else}}
			{{printf "%d %d %d" 11 11 11}}
		{{end}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `item3 := helpers.Cap(helpers.Cap(ctx.Field, "boring"), "water")`)
	hasText(t, builder, `if helpers.IsInvalid(item3) {`)
	hasText(t, builder, `helpers.AttachToNode(helpers.Printf("%d %d %d", 11, 11, 11), rootDoc)`)
	hasText(t, builder, `helpers.AttachToNode(helpers.Printf(helpers.Printf("%q", item3), "%s"), rootDoc)`)
}

func TestTextWriter_ParseTextUsingWith4(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{with cap .Items }}
			{{printf "%q" . | rewrap }}
		{{else}}
			{{printf "%d %d %d" 11 11 11}}
		{{end}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, "item3 := helpers.Cap(ctx.Items)")
	hasText(t, builder, "if helpers.IsInvalid(item3) {")
	hasText(t, builder, `helpers.AttachToNode(helpers.Printf("%d %d %d", 11, 11, 11), rootDoc)`)
	hasText(t, builder, `helpers.AttachToNode(rewrap(helpers.Printf("%q", item3)), rootDoc`)
}

func TestTextWriter_ParseTextRangeWithoutKeys(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{range .Items}}
				name item-{{.Id}} {{.PageName}} = {{.Price}}
		{{else}}
			No data to count
		{{end}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, "if helpers.Len(ctx.Items) == 0 {")
	hasText(t, builder, "for _, item3 := range ctx.Items {")
	hasText(t, builder, "No data to count")
	hasText(t, builder, "name item-")
	hasText(t, builder, "helpers.AttachToNode(item3.Id, rootDoc)")
	hasText(t, builder, "helpers.AttachToNode(item3.PageName, rootDoc)")
	hasText(t, builder, "helpers.AttachToNode(item3.Price, rootDoc)")
}

func TestTextWriter_ParseTextRangeHtmlWithoutKeys(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		<div class="item item-{{.Id}}" id="item1" event-click="shoot-starts" event-lap="{{ .Id | print }}">
			<h3 class="name item-{{.Id}}">{{.PageName}}</h3>
			<span class="price">{{.Price}}</span>
		</div>
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `attr15.Add("price")`)
	hasText(t, builder, `helpers.AttachToNode(ctx.Price, node14)`)
}

func TestTextWriter_ParseTextRangeHtmlWithoutKeys2(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{range .Items}}
		  <div class="item">
			<h3 class="name item-{{.Id}}">{{.PageName}}</h3>
			<span class="price">{{.Price}}</span>
		  </div>
		{{else}}
			No data to count
		{{end}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, "if helpers.Len(ctx.Items) == 0 {")
	hasText(t, builder, "for _, item3 := range ctx.Items {")
	hasText(t, builder, "No data to count")
}

func TestTextWriter_ParseTextRangeWithOneKey(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{range $x := .Items}}
		  <div class="item">
			<h3 class="name item-{{$x}}">{{.PageName}}</h3>
			<span class="price">{{$x.Price}}</span>
		  </div>
		{{else}}
			No data to count
		{{end}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, "if helpers.Len(ctx.Items) == 0 {")
	hasText(t, builder, "} else {")
	hasText(t, builder, "for x := range ctx.Items {")
	hasText(t, builder, "No data to count")
}

func TestTextWriter_ParseTextRangeWithKeys(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{range $x, $y := .Items}}
		  <div class="item">
			<h3 class="name item-{{$x}}">{{.PageName}}</h3>
			<span class="price">{{$y.Price}}</span>
		  </div>
		{{else}}
			No data to count
		{{end}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, "if helpers.Len(ctx.Items) == 0 {")
	hasText(t, builder, "} else {")
	hasText(t, builder, "for x, y := range ctx.Items {")
	hasText(t, builder, "No data to count")
}

func TestTextWriter_ParseTextRange(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{range .Items}}
		  <div class="item {lexus}" id="d1">
			<br/>
			<h3 class="name">{{.PageName}}</h3>
			<span class="price">${{.Price}}</span>
		  </div>
		{{end}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, "for _, item3 := range ctx.Items {")
}

func TestTextWriter_ParseJavascriptTextWith(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		 <script class="item">
			function getNames(nameList){
				console.log("Log PageName", nameList);
			}
		 </script>
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `\n\t\t\tfunction getNames(nameList){\n\t\t\t\tconsole.log(\"Log PageName\", nameList);\n\t\t\t}\n\t\t`)
}

func TestTextWriter_ParseTextMarkup(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		<footer> 
			<p>Here is the footer</p>
		</footer>
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, `domu.Element("footer", "node3")`)
	hasText(t, builder, `domu.Text("Here is the footer")`)
}

func TestTextWriter_ParseTextIfWithElse2(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{if and .User .User.Admin}}
		  You are an admin user!
		{{else}}
		  Access denied!
		{{end}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, "if helpers.And(ctx.User, ctx.User.Admin) {")
	hasText(t, builder, " Access denied!")
}

func TestTextWriter_ParseTextIfWithElse3(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{if and .User .User.Admin}}
		  You are an admin user!
		{{else if or .User .User.Admin }}
		  Access denied!
		{{end}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, "if helpers.And(ctx.User, ctx.User.Admin) {")
	hasText(t, builder, "You are an admin user!")
	hasText(t, builder, "} else if helpers.Or(ctx.User, ctx.User.Admin) {")
	hasText(t, builder, "Access denied!")
}

func TestTextWriter_ParseTextIf(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		Hello, {{if .PageName}} {{.PageName}} {{else}} there {{end}}!
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, "if ctx.PageName {")
	hasText(t, builder, "helpers.AttachToNode(ctx.PageName, rootDoc")
	hasText(t, builder, "} else {")
	hasText(t, builder, "there")
	hasText(t, builder, " there ")
}

func TestTextWriter_ParseTextWithVariableAssignmentWithOp(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{$y := (printf .Field1.Field2.Field3).Value}}
		{{$z := $y.Field1.Field2.Field3}}
		{{$ui := (cap $z.Field1 $y.Field1) }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, "y := (helpers.Printf(ctx.Field1.Field2.Field3)).Value")
	hasText(t, builder, "z := y.Field1.Field2.Field3")
	hasText(t, builder, "ui := helpers.Cap(z.Field1, y.Field1)")
}

func TestTextWriter_ParseTextWithOneVariableAssignmentWithField(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{$y := ( printf .Field1.Field2.Field3 ) }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, "y := helpers.Printf(ctx.Field1.Field2.Field3)")
}

func TestTextWriter_ParseTextWithVariableAssignmentWithField(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{$y := (printf .Field1.Field2.Field3).Value}}
		{{$z := $y.Field1.Field2.Field3}}
		{{$ui := $z }}
		{{$ui2 := $z.Field1 }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, "y := (helpers.Printf(ctx.Field1.Field2.Field3)).Value")
	hasText(t, builder, "z := y.Field1.Field2.Field3")
	hasText(t, builder, "ui := z")
	hasText(t, builder, "ui2 := z.Field1")
}

func TestTextWriter_ParseTextWithVariableUsage(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{$x := (printf .Field1.Field2.Field3 .Field1.Field2).Value}}
		{{$y := (printf $x.Field1.Field2.Field3).Value}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, "x := (helpers.Printf(ctx.Field1.Field2.Field3, ctx.Field1.Field2)).Value")
	hasText(t, builder, "y := (helpers.Printf(x.Field1.Field2.Field3)).Value")
}

func TestTextWriter_ParseTextWithMultiFunctionCallsWithNoAssignment(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{printf .Field1.Field2.Field3}}
		{{(printf .Field1.Field2.Field4)}}
		{{(printf .Field1.Field3.Field6).Value}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, "helpers.AttachToNode(helpers.Printf(ctx.Field1.Field2.Field3), rootDoc)")
	hasText(t, builder, "helpers.AttachToNode(helpers.Printf(ctx.Field1.Field2.Field4), rootDoc)")
	hasText(t, builder, "helpers.AttachToNode((helpers.Printf(ctx.Field1.Field3.Field6)).Value, rootDoc)")
}

func TestTextWriter_ParseTextBasicPipe(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ .Field | print }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, "helpers.AttachToNode(helpers.Print(ctx.Field), rootDoc)")
}

func TestTextWriter_ParseTextBasicPipe2(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{ .Field | print | cap }}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, "helpers.AttachToNode(helpers.Cap(helpers.Print(ctx.Field)), rootDoc)")
}

func TestTextWriter_ParseTextWithSingleArguments(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{$w := (printf .Field1.Field2.Field3 | printf).Value}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, "w := (helpers.Printf(helpers.Printf(ctx.Field1.Field2.Field3))).Value")
}

func TestTextWriter_ParseTextWithSplitArguments(t *testing.T) {
	var builder = textWriterTestHelper(t, `
		{{rootType string }}

		{{$x := (printf .Field1.Field2.Field3 .Field1.Field2).Value}}
	`, false)

	require.NotNil(t, builder)
	hasText(t, builder, "x := (helpers.Printf(ctx.Field1.Field2.Field3, ctx.Field1.Field2)).Value")
}

func textWriterTestHelper(t *testing.T, text string, expectErr bool) string {
	var parseBuilder, parseErr = miru.New(text, allFuncs, partialFiles)
	if expectErr {
		require.Error(t, parseErr)
		return ""
	}

	require.NoError(t, parseErr)
	require.NotEmpty(t, parseBuilder)

	var formatted, err = parseBuilder.Format("sample")
	require.NoError(t, err)
	return formatted
}

func hasText(t *testing.T, content string, match string) {
	require.True(t, strings.Contains(content, match))
}

func hasNotText(t *testing.T, content string, match string) {
	require.False(t, strings.Contains(content, match))
}

type memoryFS struct {
	files map[string]string
}

type memoryFile struct {
	name    string
	content string
}

func (m *memoryFile) Refresh() error {
	return nil
}

func (m *memoryFile) Parse(ops miru.Options, funcs map[string]string) (*miru.ParsedData, error) {
	return miru.ParseTree(m.name, m.content, ops, funcs, partialFiles, nil)
}

func (m *memoryFile) ParseFor(ops miru.Options, funcs parse.FuncMaps, t miru.Treeset) (*miru.ParsedData, error) {
	return miru.ParseTreeWithFuncMaps(m.name, m.content, ops, funcs, partialFiles, t)
}

func (m *memoryFile) ParseBlock(name string, ops miru.Options, funcs map[string]string) (*miru.ParsedData, error) {
	return miru.ParseTree(name, m.content, ops, funcs, partialFiles, nil)
}

func (m *memoryFile) ParseBlockFor(name string, ops miru.Options, funcs parse.FuncMaps, t miru.Treeset) (*miru.ParsedData, error) {
	return miru.ParseTreeWithFuncMaps(name, m.content, ops, funcs, partialFiles, t)
}

func (m *memoryFile) Read() (string, error) {
	return m.content, nil
}

func (mf *memoryFS) GetFile(name string) (miru.File, error) {
	var val, ok = mf.files[name]
	if !ok {
		return nil, nerror.New("not found")
	}
	return &memoryFile{content: val, name: name}, nil
}
