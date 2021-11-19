package stringer

import (
	"testing"

	"github.com/escaletech/go-escale/messages"
	"github.com/stretchr/testify/assert"
)

func TestJsonFromText(t *testing.T) {
	t.Run("Text block has a valid JSON", func(t *testing.T) {
		s := `\u0000\u0000\u0000\u0000\ufffd\u0002\u000csubmit\u0002\u0010checkout\u0002\u0006web\u0002ȧ\ufffd\ufffd\ufffd_\u0002\n1.1.0\u0000\u0002H6c4864dc-ff85-4d1f-b3d5-a92788c53366\u0002\ufffd\u0004https://checkout.celular\u0026fbclid=IwAR3hF_ueGiLV6\u0002.https://l.facebook.com/\u0002\ufffd\u0007{\"form_element\":\"\",\"form_classes\":\"Button__ButtonWrapper-ixybdn-0 kabpA-d FormAddress__Button-sc-17owu7c-3 laIhia\",\"form_id\":\"addressForm\",\"form_target\":\"\",\"form_url\":\"\",\"form_text\":\"Avançar\",\"email\":\"fakemail@test.com\",\"provider\":\"escale\",\"product\":\"escale-200-mega\",\"contract_type\":\"new_number\",\"cpf\":\"fakeCPF\",\"phone\":\"(27) 9999-88888\",\"segment\":\"internet\",\"utm_campaign\":\"VIV_VIV_B2CF\",\"utm_source\":\"facebook\",\"utm_medium\":\"CPC\",\"purchase_id\":\"kw4626irj\"}\u0002\u001c192.168.61.100\u0002\ufffd\ufffd\ufffd\ufffd\ufffd_\u0002\ufffd\u0003Mozilla/5.0 (Linux; Android 10; Redmi Note 8 Build/QKQ1.200114.002; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/87.0.4280.101 Mobile Safari/537.36 [FB_IAB/FB4A;FBAV/344.0.0.34.116;]\u0002\u001c192.168.43.227\u0002\u000csocial\u0002\u0010Facebook\u0000\u0002\u0010Facebook\u0002\u0006344\u0002\u00020\u0002\u000eAndroid\u0000\u0000\u0002\u0026XiaoMi Redmi Note 8\u0000\u0000\u0000\u0000\u0000\u0000\u0002Hf6f9620a-2209-450a-bd6e-16acb4667601\u00022checkout.celula.com\u0002\u003e/internet/escale-200-mega/address\u0002\u0006CPC\u0002\u0010facebook\u0002jVIV_VIV_B2CF\u0002zIwAR3hF_ueGiLV6ni5bxkGZP7SxBICOKSFt9hfmKV0peYfkJ7gZHxssX6o3jw\u0002\u0000\u0002\ufffd\u0001Button__ButtonWrapper-ixybdn-0 kabpA-d FormAddress__Button-sc-17owu7c-3 laIhia\u0002\u0016addressForm\u0002\u0000\u0002\u0000\u0002\u0010Avançar\u0002(fakemail@test.com\u0000\u0002\u001aescale-200-mega\u0002\u0010internet\u0002\u0008escale\u0002\u0014new_number\u0002\u001c128.777.887-31\u0002\u001e(27) 9999-88888\u0000\u0002@0ba8136975ed3806447ee05e74541ad3\u0000\u0000\u0000\u0002\u0010kw466irj\u0000\u0000`
		res, err := JsonFromText(s)

		assert.NoError(t, err)
		assert.Equal(t, res["email"], "fakemail@test.com")
		assert.Equal(t, res["cpf"], "fakeCPF")
		assert.Equal(t, res["phone"], "(27) 9999-88888")
	})

	t.Run("Text block doesn't have a JSON", func(t *testing.T) {
		s := "aaa"
		res, err := JsonFromText(s)

		assert.Error(t, err, messages.InputParamNotParseable)
		assert.Len(t, res, 0)
	})
}
