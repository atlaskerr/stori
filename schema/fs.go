// Code generated by "esc -o=fs.go -modtime=1546544639 -pkg=schema -include=.*\.json$ -ignore=test-fixtures/* ."; DO NOT EDIT.

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
		modtime: 1546544639,
		compressed: `
H4sIAAAAAAAC/2SQMW7DMAxFd52CUDsmcbcCXjMVKNClF1BlxmLgiALJxSh890K20hrN+vjfF8VvBwD+
mQbfg09mpe86LpgjZwuUUfTEMnYaE95CF4OFiUd/2KyN7syrcj5udNUGCRc7vrw2/6mJA2oUKkacq/yZ
SKEIFxSbob2sEGAiNeALCBZWMhZChcRqOABnsIQgOJKazKdW3WoI1fewfq7SfcGO1xEZ3v6xim0uWHdT
E8qj/xsuB/eYCiJhvocW9xvzRjatiY/zG5y368E7qbV17wX8dcVo3i3uJwAA///CisY2kQEAAA==
`,
	},

	"/oci/content-descriptor.schema.json": {
		name:    "content-descriptor.schema.json",
		local:   "oci/content-descriptor.schema.json",
		size:    2013,
		modtime: 1546544639,
		compressed: `
H4sIAAAAAAAC/7RV227bRhB991dMGaNoa1GS5fhGpCmEOEUNpHXq2C8xFGfFHZKTkrvs7CiyIujfiyUp
62LGcR/yRHJ3zsycMxfOdwAg2CUdRBBkImXU69kSTWyNKDLIrms57VGhUuxpdDFTKZaDTo1zcYaFWsN+
ctaE9WkF1KwSCfvHvfrsWQNUZnaRBBHc+C8AmDdPf1eyLZGF0AXR+oW/05SikwfnFUyJIBufyweXqcHh
UXQzDH9XYdIPT0fzo+eL3WADtNh58LrofLd8DvcHG/nsD06ekpB/jmrNluqTrZwOYVUOSCwDY4KMJiaT
goJxbsfdRu1WBoEyxory7raIbUe6yshB42MGTWc4UDwmYcUzKFCUVqK68HZp9ef1uyuYOATJEFaRgCc5
uiavTaHePiJzd77fWbSqLLMSfY5OmEz6dT3XAy4xdvwJYwm2a99a1G1NPK2YZ6XYlFWZUQxxhvE/blJA
DQebVNzrIB0gU302ZOHjC5Wnlkmy4mX0ws+bRv3yY5swVQfdqPCL75u9n36Lbva6t+Ho/uTnXyL/Pgzf
98PTX2/D0d5u0EZ3U6IV3QI1qava6BHGX09tGL5vclm9/vBs98fww213bzTvd/YHR4ve08z+X+qOvuAT
evfy9d/X55evz1ZN7EqMKSF0VVG8m6pC45mg6yxLx2padTsa6cLmEOAdOXHgLEimBBTEOaERmFKeQ6Y+
+54HvCsxFtSV/2pEq66pPcIYE8voXcboHJm0C+d13BxNKtl9FihM+Bn1PVJbdGCsQKEkzmoGDR/dgDsb
od79cXH95gz+uriCMYLwxAnqzSEs1B0VkyKI4HQwODg4HvQPjk4Onx8fH570TzYMyTSG4Tcsl8UjI5gi
P6zehPNvbJ4h5FSP0vXlGwcJ2wKmGTWkl/QKNfO8tJ2a3Cq9TY0Ei5adklgulJ/yYMK0Dnh0rbQuEsWs
ZsH6yq7NAsZ/J8So1/50dc+2DN/2+lnb/EKSV4EuXp3Dq4b12fbPeHutLXb+CwAA//9pfVCF3QcAAA==
`,
	},

	"/oci/errors.schema.json": {
		name:    "errors.schema.json",
		local:   "oci/errors.schema.json",
		size:    707,
		modtime: 1546544639,
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
		modtime: 1546544639,
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
		size:    12528,
		modtime: 1546544639,
		compressed: `
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
`,
	},

	"/oci/image-manifest.schema.json": {
		name:    "image-manifest.schema.json",
		local:   "oci/image-manifest.schema.json",
		size:    6369,
		modtime: 1546544639,
		compressed: `
H4sIAAAAAAAC/+yZW3PTuBfA3/spzj90/gPbOE0DvZBh2elQmO0MbFlo94FOoIp9HB8qS0aSm5pOvvuO
LOfi1E7cdhde9i2WdaRz10/OzQYAtDYpaPWhFRmT9Le3ZYLCl8IwEqh0R6rRNsVshNsxExSiNq22k9J+
hDFbkPyqpfDcaC4WKBYar7u/7cYeFYKJkgkqQ6hbfcg1sKNMCGmYISkWh+2bALWvKLGv7GanEWko1sig
0FQDU0MyiqkMYjQsYIZ14P101ruzj6eQagQTIcx3ApVy1B2nV7FbwoxBJd5XKVlM6dzstCe3hu0bkyVo
ddRGkRi1Su8n86fJ4oZTGTn8ir6ZysymtHwpQhot+aRhyKauk6pk431DNwtVdhK2+nBesu+2N5J6H05D
SyObUNVvF6JhNf2sI9bb3eufH3pvmBd2veeDm71nk81WhehkY+XAov9/kOa7O72S5ju9g/upPn8atOuL
5BDmoYdQKlAYokLhkxgBgyGXw6W0X5HvtaX5Mwq0cZmuK9YGJVsXkXZt4ZeLuHL6itRZ9qR1hq+yxMiR
YklEPvgR+pc6jcEtAjLMPea2bQOJ/LFwDly8YHwkFZkoftl/YXtEgMHLi3p35tl6zrzvNke3Hv/WP9/q
fPEGs5Env/Tt70PvU9d7/usXb7C12WrfrQsuOSPGgNipE1jrj3WKH3qfCk3nP//3aPP/3ucvna3BTbe9
09ubbDeb9lDDNH3HxtXy4fWfZ8cfXh/Ny0Yn6FNIqPOA2sXy6A4zg7o9Dbti47y+UJgOlMsOr0kbDVqC
iZgBBj4nFAbGxDlE7MpWGeB1gr7BIF8/bxN5xrkVYYihVGiX9FFrEqMOHLt9OYqRiWZaoFGEVxjMJAOJ
GoQ0EDPjR86Cwp6gEG6Xtvr4+8nZ2yP44+QUhghGpdpgUFX2MbumOI1bfXje6z19ut/rPt072H22v797
0D2omE6imO41mj8NMAmDI1SrI5wq3qgfHgInV6pnH95qCJWMYRxR4ZipC2KWWdsDORZcsqDafDIY13a6
UKqY2b7SShXdFm7Q7FY0NqYUy5oBjcJvKSkMbhGCq4naBlDdJmvOO0OG54qdvDqGV4UPj6pxZx1hcZah
WkOdr5kfgXX/tMfmDnHn1hCBLZy3thRL7/NyszJDphHy3YAZIBHgNXQ78DEdavyWWgucKk7MRCgglJzL
sd1UG+ZfglQBKnhMHey4THIS590BGDl94Cgeu59PvJ3BE6dQSIJxCImjzrQ1hLNMpsbtNS9UhTrlebay
JOGZ5YW85J1iRuZ9I05MBgEp9I1UmVtfji18RpS0IZYBtoGJAKSJcmONomFqUE97BgkyxPjySsAUQipm
zaLMKNXZ/zAWfjgO1xNxFVo2ossmgPkAOq5immac80PNaYzMje0pDQyW+83DCHq9J9Zx9E9C6bvQdAOg
bsbUdSFbeWpVwXX1qbUy234aZf/ToH2/E30lcTeC7h/N3fezsw7A/2PwhzH43TH8HiS+FsarY16H5P8q
la8D82Zs3vBbxEpPVUD6um9f9bReA+yrmL0S2yuO2jvA+8r+X7p2xCSOiyjsVNF/yTlz+HdE9xcq7RJj
3ZfntW0iXw+u3IIdeJNXNOnpCMhwoRb9/GxuuxnTy8RF78KCNgqdKntp8C/HTNnKjhNmaEicTAZjstXP
7W2gWDmH6yPpX2JxBbliPEXXIUhDSMgD13tsW/AjJkZYBux5WfeWPFs1XFOfebo5/1Yl15LDlz9zL1/K
NmbpU86aY8v08K78l8Ryokw2/g4AAP//gVlZF+EYAAA=
`,
	},

	"/oci/tag-list.schema.json": {
		name:    "tag-list.schema.json",
		local:   "oci/tag-list.schema.json",
		size:    432,
		modtime: 1546544639,
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
		size:    1541,
		modtime: 1546544639,
		compressed: `
H4sIAAAAAAAC/6RUS2vbQBC++1cMao92nEsp9a2EFgKhGJxbyWGtHUmTyDPq7DitKP7vZaWN65fUhhx3
Zr7HPNjfEwDI3pPPFpBVZs1iPg8mShtXYrgSLechr3Dj5gH1GXWWCxdUZtMe1+cOsI9BeNZHO7BXV9js
+mNieZeAznsyEnb1UqVBNcKQLaBwdcC+wmPIlZpYFOlXnTr06lt1MQ7embtKjM0hT9dV0lEMh6EL3PcV
wu1yXkkwdhsExx4aUQMTqCkYMggnncRgbYMRGkyJyyxldi8lmdWnmv/seMDb3Wq46QSh4lhrcBx/08hu
XaO/lIvpXDhYtgDTLZ7md5OB1+7Q04h+lqPaV6rxorOT/pfOqriHXx+uP0EEUkG5MzyawOhSzs11HmpC
tpvP/+vjtoCANgXHLSj+2GIwaBQDshGX4KAnPLQIgUpGD+sWhBGkAKsQ3NYqUYqTAeIu1ENnuZsVVCNQ
6Ioic6Tx8JOsAsdAPsashVxUMTTCPmqbdCw3stkIf4sHnKTOLb11aiNnczqwL10p3N+txkTXIjU6Hld9
wva159IoPccVPGELcajT1zY+cNZxLecHHQ+CtJvL9zOZ/bWfO3hp7CjxMCSd7Mr6EXPbfzmTfd0lG/sP
ML4f+rJTot3kTwAAAP//5ZHC3QUGAAA=
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
		_escData["/oci/content-descriptor.schema.json"],
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
