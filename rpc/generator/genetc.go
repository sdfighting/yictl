package generator

import (
	_ "embed"
	"fmt"
	"github.com/sdfighting/yictl/rpc/parser"
	"github.com/sdfighting/yictl/util"
	"github.com/sdfighting/yictl/util/format"
	"github.com/sdfighting/yictl/util/pathx"
	"github.com/sdfighting/yictl/util/stringx"
	"path/filepath"
	"strings"

	conf "github.com/sdfighting/yictl/config"
)

//go:embed etc.tpl
var etcTemplate string

// GenEtc generates the yaml configuration file of the rpc service,
// including host, port monitoring configuration items and etcd configuration
func (g *Generator) GenEtc(ctx DirContext, _ parser.Proto, cfg *conf.Config) error {
	dir := ctx.GetEtc()
	etcFilename, err := format.FileNamingFormat(cfg.NamingFormat, ctx.GetServiceName().Source())
	if err != nil {
		return err
	}

	fileName := filepath.Join(dir.Filename, fmt.Sprintf("%v.yaml", etcFilename))

	text, err := pathx.LoadTemplate(category, etcTemplateFileFile, etcTemplate)
	if err != nil {
		return err
	}

	return util.With("etc").Parse(text).SaveTo(map[string]any{
		"serviceName": strings.ToLower(stringx.From(ctx.GetServiceName().Source()).ToCamel()),
	}, fileName, false)
}
