package topo_server_test

import (
	"context"

	"configcenter/src/common"
	"configcenter/src/test/util"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var testData1 = map[string]interface{}{
	"bk_module_name": "qwq",
}

var testData2 = map[string]interface{}{
	"bk_module_name": "qwq222",
}

var tmpData = map[string]interface{}{
	"bk_module_name": "tmp_data",
}

var _ = Describe("resource pool directory test", func() {

	BeforeEach(func() {
		// 准备数据
		prepareData()
	})

	var _ = Describe("create resource pool directory test", func() {

		It("create with normal data", func() {
			data := tmpData
			rsp, err := topoServerClient.ResourceDirectory().CreateResourceDirectory(context.Background(), header, data)
			util.RegisterResponse(rsp)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
		})

		It("create with bk_module_name already exist data", func() {
			data := map[string]interface{}{"bk_module_name": testData1["bk_module_name"]}
			rsp, err := topoServerClient.ResourceDirectory().CreateResourceDirectory(context.Background(), header, data)
			util.RegisterResponse(rsp)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(common.CCErrCommDuplicateItem))
		})
	})

	var _ = Describe("delete resource pool directory test", func() {

		It("create with normal data", func() {
			id := moduleID1
			rsp, err := topoServerClient.ResourceDirectory().DeleteResourceDirectory(context.Background(), header, id)
			util.RegisterResponse(rsp)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
		})
	})

	var _ = Describe("update resource pool directory test", func() {

		It("update with normal data", func() {
			id := moduleID1
			data := map[string]interface{}{"bk_module_name": "update module name"}
			rsp, err := topoServerClient.ResourceDirectory().UpdateResourceDirectory(context.Background(), header, id, data)
			util.RegisterResponse(rsp)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
		})

		It("update with bk_module_name already exist data", func() {
			id := moduleID2
			data := map[string]interface{}{"bk_module_name": testData1["bk_module_name"]}
			rsp, err := topoServerClient.ResourceDirectory().UpdateResourceDirectory(context.Background(), header, id, data)
			util.RegisterResponse(rsp)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(false))
			Expect(rsp.Code).To(Equal(common.CCErrCommDuplicateItem))
		})
	})

	var _ = Describe("search resource pool directory test", func() {

		It("search with default query condition", func() {
			cond := make(map[string]interface{}, 0)
			rsp, err := topoServerClient.ResourceDirectory().SearchResourceDirectory(context.Background(), header, cond)
			util.RegisterResponse(rsp)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
		})

		It("search with configured condition", func() {
			queryData := map[string]interface{}{"condition": map[string]interface{}{"bk_module_name": testData2["bk_module_name"]}}
			rsp, err := topoServerClient.ResourceDirectory().SearchResourceDirectory(context.Background(), header, queryData)
			util.RegisterResponse(rsp)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			Expect(rsp.Data.Count).To(Equal(int64(1)))
			Expect(rsp.Data.Info[0].String("bk_module_name")).To(Equal(testData2["bk_module_name"]))
		})

		It("search with configured exact is true", func() {
			queryData := map[string]interface{}{"exact": true, "condition": map[string]interface{}{"bk_module_name": "qq"}}
			rsp, err := topoServerClient.ResourceDirectory().SearchResourceDirectory(context.Background(), header, queryData)
			util.RegisterResponse(rsp)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			Expect(rsp.Data.Count).To(Equal(int64(0)))

			queryData = map[string]interface{}{"exact": true, "condition": map[string]interface{}{"bk_module_name": testData1["bk_module_name"]}}
			rsp, err = topoServerClient.ResourceDirectory().SearchResourceDirectory(context.Background(), header, queryData)
			util.RegisterResponse(rsp)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			Expect(rsp.Data.Count).To(Equal(int64(1)))
			Expect(rsp.Data.Info[0].String("bk_module_name")).To(Equal(testData1["bk_module_name"]))
		})

		It("search with configured exact is false", func() {
			queryData := map[string]interface{}{"exact": false, "condition": map[string]interface{}{"bk_module_name": "qwq"}}
			rsp, err := topoServerClient.ResourceDirectory().SearchResourceDirectory(context.Background(), header, queryData)
			util.RegisterResponse(rsp)
			Expect(err).NotTo(HaveOccurred())
			Expect(rsp.Result).To(Equal(true))
			Expect(rsp.Data.Count).To(Equal(int64(2)))
		})
	})

})

var moduleID1, moduleID2 int64

func prepareData() {
	if moduleID1 != 0 {
		id := moduleID1
		rsp, err := topoServerClient.ResourceDirectory().DeleteResourceDirectory(context.Background(), header, id)
		util.RegisterResponse(rsp)
		Expect(err).NotTo(HaveOccurred())
		Expect(rsp.Result).To(Equal(true))
	}
	if moduleID2 != 0 {
		id := moduleID2
		rsp, err := topoServerClient.ResourceDirectory().DeleteResourceDirectory(context.Background(), header, id)
		util.RegisterResponse(rsp)
		Expect(err).NotTo(HaveOccurred())
		Expect(rsp.Result).To(Equal(true))
	}

	rsp, err := topoServerClient.ResourceDirectory().CreateResourceDirectory(context.Background(), header, testData1)
	util.RegisterResponse(rsp)
	Expect(err).NotTo(HaveOccurred())
	Expect(rsp.Result).To(Equal(true))

	moduleID1 = int64(rsp.Data.Created.ID)

	result, err := topoServerClient.ResourceDirectory().CreateResourceDirectory(context.Background(), header, testData2)
	util.RegisterResponse(rsp)
	Expect(err).NotTo(HaveOccurred())
	Expect(rsp.Result).To(Equal(true))

	moduleID2 = int64(result.Data.Created.ID)
}
