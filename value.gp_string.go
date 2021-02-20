///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/gxlb/gogp]
// Last update at: [Sat Feb 20 2021 11:08 CST]
// Generate from:
//   [github.com/gxlb/cli/internal/gp/value.gp]
//   [github.com/gxlb/cli/value.gpg] [value_string]
//
// Tool [github.com/gxlb/gogp] info:
// CopyRight 2021 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Site    : https://github.com/vipally
// Version : 3.0.0.final
// 
///////////////////////////////////////////////////////////////////

package cli

import (
	"flag"
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////

// StringValue define a value of type string
type StringValue struct {
	Target      *string  // Target value pointer outside
	Default     string   // Default value
	DefaultText string   // Default value in help info
	Enums       []string // Enumeration of valid values
	Ranges      []string // {[min,max),[min,max),[min...)} ranges of valid values
	value       string   // The value from ENV of files
	hasBeenSet  bool
}

func (v *StringValue) IsSet() bool {
	return v.hasBeenSet
}

func (v *StringValue) Apply(f *Flag, set *flag.FlagSet) error {
	if l := len(v.Enums); l > maxSliceLen {
		return fmt.Errorf("flag %s.Enums too long: %d/%d", f.logicName, l, maxSliceLen)
	}
	if l := len(v.Ranges); l > maxSliceLen {
		return fmt.Errorf("flag %s.Ranges too long: %d/%d", f.logicName, l, maxSliceLen)
	}

	return nil
}

func (v *StringValue) String() string {
	return ""
}

// check if value if valid for this flag
func (v *StringValue) validValue(f *Flag, value string) error {
	if len(v.Enums) > 0 {
		found := false
		for _, v := range v.Enums {
			if value == v {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("flag %s value %v out of Enums: %v", f.logicName, value, v.Enums)
		}
	}
	if len(v.Ranges) > 0 {
		found := false
		for i := 0; i < len(v.Ranges); i++ {
			min := v.Ranges[i]
			max := min
			if i++; i < len(v.Ranges) {
				max = v.Ranges[i]
			}
			if value >= min && value < max {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("flag %s value %v out of Ranges: %v", f.logicName, value, v.Enums)
		}
	}
	return nil
}
