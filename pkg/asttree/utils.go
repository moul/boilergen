package asttree

import (
	"fmt"
	"go/ast"
)

func resolveFieldTypes(t ast.Expr, pkgName string) string {
	switch t1 := t.(type) {
	case *ast.StructType:
		return "struct{}"
	case *ast.InterfaceType:
		return "interface{}"
	case *ast.SelectorExpr:
		// we have to override this crap...
		return fmt.Sprintf("%s.%s", t1.X, resolveFieldTypes(t1.Sel, ""))
	case *ast.StarExpr:
		return fmt.Sprintf("*%s", resolveFieldTypes(t1.X, pkgName))
	case *ast.Ident:
		if pkgName != "" {
			return fmt.Sprintf("%s.%s", pkgName, t1)
		}
		return fmt.Sprintf("%s", t1)
	case *ast.MapType:
		return fmt.Sprintf("map[%s]%s", resolveFieldTypes(t1.Key, pkgName), resolveFieldTypes(t1.Value, pkgName))
	case *ast.ArrayType:
		l := ""
		if t1.Len != nil {
			// we have an array, not a slice.. pity...
			l = fmt.Sprintf("%s", t1.Len)
		}
		return fmt.Sprintf("[%s]%s", l, resolveFieldTypes(t1.Elt, pkgName))
	default:
		return fmt.Sprintf("UKNOWN: +V", t)
	}
}
