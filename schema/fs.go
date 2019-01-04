// Code generated by "esc -o=fs.go -pkg=schema -include=.*\.json$ -ignore=test-fixtures/* ."; DO NOT EDIT.

package schema

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
	"time"
)

type _escLocalFS struct{}

var _escLocal _escLocalFS

type _escStaticFS struct{}

var _escStatic _escStaticFS

type _escDirectory struct {
	fs   http.FileSystem
	name string
}

type _escFile struct {
	compressed string
	size       int64
	modtime    int64
	local      string
	isDir      bool

	once sync.Once
	data []byte
	name string
}

func (_escLocalFS) Open(name string) (http.File, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	return os.Open(f.local)
}

func (_escStaticFS) prepare(name string) (*_escFile, error) {
	f, present := _escData[path.Clean(name)]
	if !present {
		return nil, os.ErrNotExist
	}
	var err error
	f.once.Do(func() {
		f.name = path.Base(name)
		if f.size == 0 {
			return
		}
		var gr *gzip.Reader
		b64 := base64.NewDecoder(base64.StdEncoding, bytes.NewBufferString(f.compressed))
		gr, err = gzip.NewReader(b64)
		if err != nil {
			return
		}
		f.data, err = ioutil.ReadAll(gr)
	})
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (fs _escStaticFS) Open(name string) (http.File, error) {
	f, err := fs.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.File()
}

func (dir _escDirectory) Open(name string) (http.File, error) {
	return dir.fs.Open(dir.name + name)
}

func (f *_escFile) File() (http.File, error) {
	type httpFile struct {
		*bytes.Reader
		*_escFile
	}
	return &httpFile{
		Reader:   bytes.NewReader(f.data),
		_escFile: f,
	}, nil
}

func (f *_escFile) Close() error {
	return nil
}

func (f *_escFile) Readdir(count int) ([]os.FileInfo, error) {
	if !f.isDir {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is not directory", f.name)
	}

	fis, ok := _escDirs[f.local]
	if !ok {
		return nil, fmt.Errorf(" escFile.Readdir: '%s' is directory, but we have no info about content of this dir, local=%s", f.name, f.local)
	}
	limit := count
	if count <= 0 || limit > len(fis) {
		limit = len(fis)
	}

	if len(fis) == 0 && count > 0 {
		return nil, io.EOF
	}

	return []os.FileInfo(fis[0:limit]), nil
}

func (f *_escFile) Stat() (os.FileInfo, error) {
	return f, nil
}

func (f *_escFile) Name() string {
	return f.name
}

func (f *_escFile) Size() int64 {
	return f.size
}

func (f *_escFile) Mode() os.FileMode {
	return 0
}

func (f *_escFile) ModTime() time.Time {
	return time.Unix(f.modtime, 0)
}

func (f *_escFile) IsDir() bool {
	return f.isDir
}

func (f *_escFile) Sys() interface{} {
	return f
}

// FS returns a http.Filesystem for the embedded assets. If useLocal is true,
// the filesystem's contents are instead used.
func FS(useLocal bool) http.FileSystem {
	if useLocal {
		return _escLocal
	}
	return _escStatic
}

// Dir returns a http.Filesystem for the embedded assets on a given prefix dir.
// If useLocal is true, the filesystem's contents are instead used.
func Dir(useLocal bool, name string) http.FileSystem {
	if useLocal {
		return _escDirectory{fs: _escLocal, name: name}
	}
	return _escDirectory{fs: _escStatic, name: name}
}

// FSByte returns the named file from the embedded assets. If useLocal is
// true, the filesystem's contents are instead used.
func FSByte(useLocal bool, name string) ([]byte, error) {
	if useLocal {
		f, err := _escLocal.Open(name)
		if err != nil {
			return nil, err
		}
		b, err := ioutil.ReadAll(f)
		_ = f.Close()
		return b, err
	}
	f, err := _escStatic.prepare(name)
	if err != nil {
		return nil, err
	}
	return f.data, nil
}

// FSMustByte is the same as FSByte, but panics if name is not present.
func FSMustByte(useLocal bool, name string) []byte {
	b, err := FSByte(useLocal, name)
	if err != nil {
		panic(err)
	}
	return b
}

// FSString is the string version of FSByte.
func FSString(useLocal bool, name string) (string, error) {
	b, err := FSByte(useLocal, name)
	return string(b), err
}

// FSMustString is the string version of FSMustByte.
func FSMustString(useLocal bool, name string) string {
	return string(FSMustByte(useLocal, name))
}

var _escData = map[string]*_escFile{

	"/oci/catalog.schema.json": {
		name:    "catalog.schema.json",
		local:   "oci/catalog.schema.json",
		size:    401,
		modtime: 1546544737,
		compressed: `
H4sIAAAAAAAC/2SQMW7DMAxFd52CUDsmcbcCXjMVKNClF1BlxmLgiALJxSh890K20hrN+vjfF8VvBwD+
mQbfg09mpe86LpgjZwuUUfTEMnYaE95CF4OFiUd/2KyN7syrcj5udNUGCRc7vrw2/6mJA2oUKkacq/yZ
SKEIFxSbob2sEGAiNeALCBZWMhZChcRqOABnsIQgOJKazKdW3WoI1fewfq7SfcGO1xEZ3v6xim0uWHdT
E8qj/xsuB/eYCiJhvocW9xvzRjatiY/zG5y368E7qbV17wX8dcVo3i3uJwAA///CisY2kQEAAA==
`,
	},

	"/oci/errors.schema.json": {
		name:    "errors.schema.json",
		local:   "oci/errors.schema.json",
		size:    707,
		modtime: 1546544737,
		compressed: `
H4sIAAAAAAAC/6yRzU7DMBCE732KleHYNtyQekUcOCEhXsDE22arxGt2N5Ui1HdHTkyJws8BcZ2Zb0b2
vq0AwF1TcDtwjVnaVRUnjDVH8xRRdMtyqLRusPMVirC49cRM2ow7KsfNpI5QEL+3zc1toa8KGFBroWTE
McPPDSkk4YRiA5RdBQ8tqQHvYRxV6FDVH1BB0HqJGGAv3IE1CIpyQoHQIxiDj0Dx5FsKIPjao9q2LJcV
QnU7GF+e1al/pmSRDLuF9mPDp11zwO+M7NmQsufUhOLBLSPn9Ze2gOap/b++8oN/LVz9Un+B+eWItc3g
efCS8iJ++AiNxVPMGVk7Jh7vHuA+HwaeUBNHxXLD5dB59R4AAP//RJkb88MCAAA=
`,
	},

	"/oci/image-config.schema.json": {
		name:    "image-config.schema.json",
		local:   "oci/image-config.schema.json",
		size:    7271,
		modtime: 1546544737,
		compressed: `
H4sIAAAAAAAC/9xZS2/cyBG+61cUJgFyyFiTwEiCCMhBkL3aQeS1oYdzWARmkywOe9XsZqqLkomF/ntQ
3U1qKHGkkePsGquLZvpZ/dVXz/n5AAAWv9fl4ggWNXN7tFq5Fm3hLCttkfyho81KN2qDq0ZZXaHnxTLu
8kWNjdra+ZN39lUcDdtKUhW/+tPfVnHsd2ljS65FYo1+cQRBAhlVVNSaseCOcGtcpkr0BemWtbNy22WN
cPLhCrZ3wG2tixq4Rsi1VaTRg7bhexAeFCHknTYM7IA6C84ewomzld50pORoDxffv786ewOdxyUoW4Ju
WoMNWn4wb0skz8qWS7hRpkMPRnvGcrjx1MGZsptO7i1d0Y1HQOUIstP3x+cn32eHEY70Ru5befbCM2m7
WaSZu+UIT8e1o6eBOdU36IMEVjUoT1g5AmyUNqDKktB7cFVY0CJ5Z0GmLWvuE34FoZKHhNd7IPSts17n
BoPojdKBF9pu7qF92TuKAPnzCsbPWHQBs1aRapCRfBIyqSFH0VQJyoOCXHnhAFrRbZBPwUhi6PxUYris
tYdKoymhUFaOsp0xS1FgAkLOU7afl8PXrjOlbPMtFrrSIgZH9GRpAnkUYArRLP/T3ElTPhqcQegNVqoz
DIo2gV1eWC1XomXqW6ctPxYCLmv0OFBWFSzQlfEkH1TeqF4eRdgaVWAJeR8wuH9kQDg+cwrxIawrUBay
t6MAWbxJeGQd3x+yFLGioVSaPEeZB3Gzk6bMQBGpfkvR2jJSS8hR3zwSRAk1o0VPQE6wacbGzwG6m6rj
393j44YdQbzphgerF/c47KPP4+BBBIOJRjshoU9qbBpREbv08kT3iYrBsyL2X6rpbfbso/NvDPCbfZAW
vUhwkHiQ3HXlqFEB/Ozj8fkPx+/e/uPj8fnH47Ort9keUMpJDdJGoNJc720x3xZ6n1vnsfzgiP1+hPUY
IGtlR2SlnAAVuWbiewd/zOJyk/tds4dr7P3gfB8rQo5dcdFmy/S5K+WzhE/5mkWoZVPSBbTk2BXOQI5y
XSabQVdT3/PLqrNVzEj2w25/nxYe/vzn5d2OyW1FuvwnLHgxs+ruBXSYO+UhH85UjmYvJlwGPx4iacRC
bCvXTIp6aJBVqViF7OFRNNIeUizs4d3VxWXwd7JKWetSxkSdQf+LgLvDrv4P4F6way/0xirzhQALRr73
jA0UyhjgWjHcamNCSoKWh3Tg3gyDhWoOBgA+XD5Ynxq+h6xxaovZxfpUPGK2DCrUVvLeAsP4P9dnZ1kw
yYv16fnlu/UPf3ydLXZDMwfwQ2iuPNK+oHQeKYjsCK7Wb1LmpiUdbI1iecGrZLkFeKYuVgoBLGWMu/Uw
TgtQ5Ay4G6R0kBwfc2VyhaTOUjUon4irCvEXctfggGK2k6L2GJvnc6CnEqnvHMGZtt3nkNSWSdF+KTIP
WVLlRP6wlcIdujyCTCQWh9npMvzzSEcbcl2bBo82aUI+juNhlS6zkMJlcXyVych86jY+OCwNPtN37VAp
UR/Hx0pjo2/QJtlWQTShWLZCLlat8v62zGLQmPJV3qXa1ojT/h859dGZrsGXhbVSExbsQqoQ1+SC9m2N
hBNSaCn+rtH0ovlb0owQHN7ILHaTcDhY0G8sXvzL0bW2mzd6L+u9QE5ZbUck7uo2bh9hH8uBrYx0RNzO
RpJoaI/Ncle2O3PnF4b5J7l4MAvZDmy3yuRYiD9dJx9LVZBri6VwLhT8wFoqf97qhxjVi09Tfqjul8m5
p2KqxCockfdw/t0JvH79+u9L8FiE8PuXw79Oi9e4VW6XK1/JdS+q/2vtBeunH/YmWlxqZqQtQglURR0f
dAhrC45KpOg9YinJDozyPBV5Pqt+qgzf1XB5MhzFDQNx806bEgJzZzjzZfm81IA4X1DOs6Pz7Bpo0Hu1
weDapsQe6fH1RJzj7bdE4L2Z/BVQ+JT3L+DPUOFPO3FfX0XYtNx/CmfuLd3YLdM+Nt3YQaPoWgqsbRPV
ISUdmohQaYMpTy11VUnhJycIESVBJcmMZg4oHXr7B4bCUexAhvuUFe/eKZMIkeIAOceVH/T9DEq5cwaV
fQzTk8FtZyCcdemTavve8Tn/fNMzZrMREXFO0UoTgvcGETva2v+qDe2hl/3+4oWd7Kiw58FIir3GHggr
JLRFCghR/xKTRY7U2cbEzLx/1Oht1HXaGZGLPWiola+hxBZtCW5oSI6Eldn9+7ZC70+63C/JtKm/6aoH
T5E74w8Xe0W2X71rlJY+++JQ3Oc4GH4WXu2l4nhAyLBwg1aoj6EjSSRlZzCJHtAWrrMs5YGFzl5bd2tT
4ndba4Nwg6SrXozGEXS2VUVI8ZSd+alicIe2axZH8OMceFHOR6D9++vkgIT/6TSFYPnjDjbN4X0wK8eu
jPJgvHXuvumvbtu+6qG5Hoz3LVizCVe9P1nDOlhU9Dvp972Hktwd/DcAAP///hoCxGccAAA=
`,
	},

	"/oci/image-index.schema.json": {
		name:    "image-index.schema.json",
		local:   "oci/image-index.schema.json",
		size:    11945,
		modtime: 1546565443,
		compressed: `
H4sIAAAAAAAC/+xaa1PbPBb+zq84m3Z2dpfEhEu5Tbc7GWBbZmjDculO26GNYh/bKrbklWRC6PDfdyTb
udV2RIH2pW++JdLRkfToOY8k63xbAoDGc+o1dqERKpXsrqzwBJnLmSKUoZAOF8EKjUmAK5R5eN1oZk2k
G2JMJpp9lZy1slLTxhPEV6321kpW9ixvmAieoFAUZWMXTPe6lDDGFVGUs8liXeOhdAVNdJXu7CykEnIf
Q8iHKYGIPlWCiCHEqIhHFHHguLB6e356BqlEUCHCuCcQaYTSycaV95YQpVCw47JB5ibOt9Xm7XfFukYN
E9RjlEpQFjSm6m/H/24nOyza8P5XdFXRZmTSiAmjPkplA8vJwX/OD08O9svwgYhKBdyHkUPwuQCZoEt9
6kISEeVzEUsH/hvSSGM1CbSBsI+QCJTIVNNAKekNapcGViHIEN52PmirGxR8ClaqMC6BcmYSHSgKuDCj
E+ijQOZSFgCBfsT7U6tVSadxdSWvfhHB7sI0C8rZca+MhKVsnEfLmkYNjwYolSXIGidXDBPFA0GSkLrg
huheyjSGzE9Bq6z/JlBm/uagQe8liQIuqArjV7svtVx56L3q1SKt+/38ibRu2q2di+W//Wv307LzpXUx
Kvn7P3b1707rY7u1888vrYvl541abCqgLsEmRo+Ss6yZDTwW8+i0PuYDH//8y7Pnf219/uIsX3xrN1fX
Nm9X7MwebJ6FiFRNk7Bh12/swqdyetbQO7EIk6IT4YZUoatSgfOMtT2yNK4e04zruDHH7KKuvjTWxu65
fOjxMk9w6t1vzEt3rKma5B9+dde3NxtNG1Rjb3PDzlLEtnabG0+KWR4RA8oWxLILQ0OYJ7W8ggSc+dFw
scK/UDqeEmN8gdiXi63maWw1NpZJ4t7FMkIr25gm0tJtZmrpV67vtK+fVMRElKXXi3hZ6KsVAgzVQl4X
dLFFgCfIFnx5HL48JR4kEWE7Cxb8ptc0ySMiqFys7588ygeUeXzws3lQUnxR/uHa5o3DeuEtn77yly2U
5uFg7/gcJntw4DAmAYJ50kQJp2+650f7kEpsAmEe0DiJMEaWvx2N6pmHQirCvCZckShFad7V0CseKF5z
OCIsSLVvj7updmGesnqvu52TvTe9GgZarfh8FtucauZfF20uijZXRLvLodW10PJCOO8qeNG87ytahYOG
j0RzS96dviPWGiqfHpy8P9g3tPFT7RGuUEjDw/xZrHi2Nex06jhV/vL6Q/OuF7aRD/MSfFfo+F1BQ2Ak
Hr0+a/SIoiwAOZQKYxiE1A1NFc3iXEI/pZECxUGkDDhzYI8znwapmAnxx5SA7mnPubcCFK869VGQf6Gf
YzT6zltvV3zdq7fKvmjU2+T32DkbZH57maMs5mw7RwvyE1K91dz98/E0g0vnx2Wje3x22H3XOSrb9QjL
czK4D9nQZBOQuGFuM8ySKmLCPKK4GEL3FPKR/B6C4uSi+bCw6gjPHVeKjyIiQK0I/SGocJTFogsi3tdH
j2lledv5oI1MAgk3eSTjJJ1BiAJhPButZIwruGR8wLT5gItLGFCVqV3IpdIrmVs78J5E1CuEigickbWW
hz5l6DmAgQO91bbTdlY31nfWndX25mYPOIM8OGpJcb8YuCKCEqYeYaEyx8VC7R2fP8AsrHNoBP4vpQK9
akGfPvTW7I3WR21FVWTm0N07zI+4h+OsvaoJ22f5SHqDd0mkmncu1/5Mdk9/qFA2i5USZGBSr5ApB6ZP
SHhNdVxIDiokCgi4EdU77IBGEYTkCrXy4XWCro5Ak6KmN1+TcZR5hD76XKB26aKUlAUOHGb9RsgCFY5G
gUpQvEJv1NLjmIVfTFR+vijm4+WNm1Nd5YeGd12TPadEqk8KFRlhMbmmsdn/d9bW1te31trrm9svNra2
Xmy3t8tbUJa3aNk2KZacMoUBCps1T0VkmzzXGaUZnp8cSfAFnzyKFaDEZKjR8PiARZx4lYDM2WoaPhcx
0bLRSAWtDJ/5UV2fAle+/UzLwIyH2tDPYqguSawyuW6pRgCmQ38vh3p/lFA5mzhZGful6alTIExkp5Yn
ttXls/2kNLaKZR8PPctKfl92RvgxJTP+xhvvv43oUDlxWvju5tbMLIrk2t5aT+/oyKS+8/WJezkgQotP
nBBF+zSiapht9TzyUEzdCve5e4lCiyVmu30mYlSCTzHyMnnUyuWGhAXTR7wJ5VmbKh7Jy1oZsjMSYiIi
w7eM/zOAlyQ4L41I/R05b5f+HwAA///UmD65qS4AAA==
`,
	},

	"/oci/image-manifest.schema.json": {
		name:    "image-manifest.schema.json",
		local:   "oci/image-manifest.schema.json",
		size:    5233,
		modtime: 1546544737,
		compressed: `
H4sIAAAAAAAC/+yY33PTRhDH3/krtobpQONfmELAQ+lkCJ1mBgqFpA9kDDlLK2nJ6U7creIIJv9753SS
HTuyrUDbvPTNOt3e7e3ud+9jfb0FAJ07FHbG0EmYs/FgoDNUgVYsSKGxfW3iAaUixkEqFEVoudP1VjZI
MBWXLD9ZrXp+tDQLjYi4N9wd+LHblWFmdIaGCW1nDKUHblQopVkwaXV52L0J0QaGMvfKbXaYkIVqjQIq
Ty0IMyU2whSQIotQsOjDm3rWq6N3h5BbBE4QFjuBySXavver2i0TzGjUmyYnqyn9r/e7F1eG3RsuMnQ+
Wjak4s7S+4vF08XlDWsbPf2EAdc28ymdQKuI4s0x2YN6QBuItAGDERpUAakYBEylnq6ccsPx1mbiJvLR
OivbctMiQ1fydDVbG3PWOL0TUuxE0yKSLhiBKTLWsRFZQgEECQanNk/BLwI6KiPmt+0CqfKxCg6cPBUy
1oY4SZ+NnzoVhxg+O1kfTrfph2PR+zLsPZns3P11fLzT/9ibzEfu/TR2v/d674e9J7987E127nS61yv6
lWCkGJI49AZb47HN8b3e+8rTxc8fbt/5sffhY39n8nXYvT96dDFoN+17D2bpC7ZWy9sXfx4dvH2xv5CN
zTCgiNCWCXWLldmdFoy2W6fdiFmpL1Tch2XZ4TlZtmA1cCIYBASSUDHMSEpIxJlTGeB5hgFjWK5ftomy
4vyKMMVIG3RLBmgtqbgPB35fiSrmZO4FsiE8w3BuGWq0oDRDKjhI/Amq84SVcXdpq3e/vz56uQ9/vD6E
KQKb3DKGTbJPxTmledoZw5PR6MGD3dHwwaPHD3/e3X34ePi4YTqpanqv1fw6waQYYzSbM5wb2aof7oEk
L9Wjty8tREanMEuoCkwdglQU7uyhnimpRdh8fGJM13a6SJtUuL7SyQ1dNW7R7DY0NmGMKNrdXwY/52TQ
IcRxgybWNoDmNrkYnCxdksSydOz18wN4XsVwf37nda5zoUpRoNkCGS9EkIALf91jy4D4e2uKIC7dt06K
S+9LuTmbqbAI5W4gGEiFeA7DPrzLpxY/5+4E3hVvxgkqiLSUeuY2tSyCU9AmRAN3qY99X0ne4ng4Adb1
g0R11/+817s/uecdikgJCRFJtIV1B5Gi0Dn7vRZCNWhzWVaryDJZOF4oJe8dY132jTTjAkIyGLA2hV9f
zxweJpR1IdUhdkGoEDQn5WHZ0DRntHXPIEVMQq6uBMIg5GreLJYZpbn6vxd9ttBPCwC6IQa6Dga1IKF2
MNTEQ42dYxsVNbeb9Wx0k3j0TxPSt7XijajUipb+a2D6tnOuI6f/4en74On6/PQNCLWVoppzvo6l/lWc
2kZU7aCq5Z/IjZFqoKsrjXYVP9di1hrS2gRbjby1glzXpK6N/b/xe8dSEBZ05r8V/YXG+gLY9iVoazso
14Mzv2AffiuVS7YeAR1d0lxQ3sFdP6OmvZPRiSMhVDY3juqC05kwTsFpJpimJIkLmJFTuXS4Vq1c0s++
Dk6xYsQzIXP0nYAsRIQy9D3GyT9IhIpxmYAW8h0tDc81OmoK7IoOy7Ly8W0qopWAr352WqXmW/MyWa6O
g1TECK+WPxGuFsTFrb8DAAD//ziD9NVxFAAA
`,
	},

	"/oci/tag-list.schema.json": {
		name:    "tag-list.schema.json",
		local:   "oci/tag-list.schema.json",
		size:    432,
		modtime: 1546544737,
		compressed: `
H4sIAAAAAAAC/3SQwWrrQAxF9/4KMe8tk0x3BW+7KhS6yQ+otmIr2KNB0saU/HsZe5KaQJc+vudeNN8N
AIT/3IcWwuie2xglU+okOXIitZPoEK0bacboOFic2DwcNm/jO/dqko4bXcVe8eLHl9fa8K+KPVmnnJ0l
Ffk8skFWyaS+QN02QChbIBfAaYIyDpwAQSmLsYsup9pXXSYLLaw3FZpwpt13Qb7kgoK5chpC/XM73JUy
8qSw0/zE/m7al+1TqIrLY655xIKzT2vi8+0dzuXEj9/3vdvydaXOQ3NrfgIAAP//IdRifbABAAA=
`,
	},

	"/stori/server-config.schema.json": {
		name:    "server-config.schema.json",
		local:   "stori/server-config.schema.json",
		size:    729,
		modtime: 1546544639,
		compressed: `
H4sIAAAAAAAC/5SRwU4CQQyG7/sUzegRWC7GyAOYmHggwReoM93dwjIz6VQiMby7GXYkKKuGa/t/f/8/
/agAwNyyMwswnWpc1HXSILzFltIsSFsn29EW60SyI5na4BtuzWTght0Zu07BT4fpEXaCjU7n98XlpoCO
khWOysFneHX0hsH7TTDPwaHirOijhEiiTMks4Jg5T9E5oXQ+GvF+6QielnUXknrcEqB3EIMoaICek5KH
4Mud4qD7SBlNKuxbUzaHL4nR/r+bz6vfyxRktFLZWRJ95J4uNiO3lqhd7vJ+N3+ADHLDFpW+Xfuz2I96
Rbuh/bURovAOlWBDe2i4p8m1EarRNCcmvK7J6ukf1Ul3oThUnwEAAP//T/d9VtkCAAA=
`,
	},

	"/": {
		name:  "/",
		local: `.`,
		isDir: true,
	},

	"/oci": {
		name:  "oci",
		local: `oci`,
		isDir: true,
	},

	"/stori": {
		name:  "stori",
		local: `stori`,
		isDir: true,
	},
}

var _escDirs = map[string][]os.FileInfo{

	".": {},

	"oci": {
		_escData["/oci/catalog.schema.json"],
		_escData["/oci/errors.schema.json"],
		_escData["/oci/image-config.schema.json"],
		_escData["/oci/image-index.schema.json"],
		_escData["/oci/image-manifest.schema.json"],
		_escData["/oci/tag-list.schema.json"],
	},

	"stori": {
		_escData["/stori/server-config.schema.json"],
	},
}
