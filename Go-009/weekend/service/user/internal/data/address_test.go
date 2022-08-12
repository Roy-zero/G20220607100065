package data_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"user/internal/biz"
	"user/internal/domain"
	"user/internal/testdata"
)

var _ = Describe("UserAddresses", func() {
	//var ro biz.UserRepo
	//var uD *biz.User
	var ao biz.AddressRepo
	var aD *domain.Address

	BeforeEach(func() {
		ao = NewAddressRepo(Db, nil)
		aD = testdata.Address()
	})

	It("CreateAddress", func() {
		a, err := ao.CreateAddress(ctx, aD)
		Ω(err).ShouldNot(HaveOccurred())
		// 组装的数据 mobile 为 17679215805
		Ω(a.Mobile).Should(Equal("17679215805")) // 手机号应该为创建的时候写入的手机号
	})

})
