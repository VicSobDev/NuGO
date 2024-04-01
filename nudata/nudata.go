package nudata

import (
	"math"
	"math/rand"
	"time"
)

type Jvq struct {
	JvqtrgQngn JvqtrgQngn `json:"jvqtrgQngn"`
	Jg         string     `json:"jg"`
}

type JvqtrgQngn struct {
	Oq    string `json:"oq"`
	WFI   string `json:"wfi"`
	Ji    string `json:"ji"`
	Oc    string `json:"oc"`
	Fe    string `json:"fe"`
	Qvqgm string `json:"qvqgm"`
	Jxe   int64  `json:"jxe"`
	Syi   string `json:"syi"`
	Si    string `json:"si"`
	Sn    string `json:"sn"`
	Us    string `json:"us"`
	Cy    string `json:"cy"`
	Sg    string `json:"sg"`
	SP    string `json:"sp"`
	Sf    string `json:"sf"`
	Jt    string `json:"jt"`
	Sz    string `json:"sz"`
	Vce   string `json:"vce"`
	Vp    string `json:"vp"`
	NS    string `json:"ns"`
	Qvg   string `json:"qvg"`
}

func Gen() string {

	data := Jvq{}

	data.JvqtrgQngn.Oq = "562:150:1921:897:1921:897"
	data.JvqtrgQngn.WFI = "flap-1"
	data.JvqtrgQngn.Ji = "2.3.1"
	data.JvqtrgQngn.Oc = "4o24q28n13rpsrs2"
	data.JvqtrgQngn.Fe = "1921k897 24"
	data.JvqtrgQngn.Qvqgm = "0"
	data.JvqtrgQngn.Jxe = genJXE()
	data.JvqtrgQngn.Si = "si,btt,zc4,jroz"
	data.JvqtrgQngn.Sn = "sn,zcrt,btt,jni"
	data.JvqtrgQngn.Us = "6102r54820166q8o"
	data.JvqtrgQngn.Cy = "ZnpVagry"
	data.JvqtrgQngn.Sg = `{\"zgc\":0,\"gf\":snyfr,\"gr\":snyfr}`
	data.JvqtrgQngn.SP = `{\"gp\":gehr,\"ap\":gehr}`
	data.JvqtrgQngn.Sf = "gehr"
	data.JvqtrgQngn.Jt = "31q446612o858870"
	data.JvqtrgQngn.Sz = "12r0213093p022r6"

	return ""
}

func genJXE() int64 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	randomNumber := math.Floor(1e6*rand.Float64()) + 1e3

	return int64(randomNumber)
}
