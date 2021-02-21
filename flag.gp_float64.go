///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/gxlb/gogp]
// Last update at: [Sun Feb 21 2021 11:15 CST]
// Generate from:
//   [github.com/gxlb/cli/internal/gp/flag.gp]
//   [github.com/gxlb/cli/flag.gpg] [flag_float64]
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

// Float64Flag define a value of type float64
type Float64Flag struct {
	//
	//name related area
	//
	LogicName string   // logic name of the flag
	Name      string   // name of the flag
	Aliases   []string // aliases of the flag
	Usage     string   // usage string
	Required  bool     // if required
	Hidden    bool     // hidden this flag
	EnvVars   []string // environment values
	FilePath  string   // file path
	//
	//value related area
	//
	Target      *float64  // Target value pointer outside
	Default     float64   // Default value
	DefaultText string    // Default value in help info
	Enums       []float64 // Enumeration of valid values
	Ranges      []float64 // {[min,max),[min,)...} ranges of valid values
	//
	////////////////////////////////////////////////////////////////////////////
	//area for parsing
	target *float64 // target value pointer(maybe new(float64) if Target not set)
	info   FlagInfo // parsed info of this flag
}

// Init verify and init the value by ower flag
func (v *Float64Flag) Init(namegen *NameGenenerator) error {
	v.info.Flag = v
	v.info.EnvVars = v.EnvVars
	v.info.Usage = v.Usage
	v.info.DefaultText = v.DefaultText
	v.info.Required = v.Required
	v.info.Hidden = v.Hidden
	v.info.FilePath = v.FilePath
	v.info.LogicName = logicName(v.LogicName)
	v.info.Name = namegen.GetOrGenName(v.Name)
	v.info.HasBeenSet = false
	v.info.DispName = v.Name
	mergeNames(v.Name, v.Aliases, &v.info.Names)

	if l := len(v.Enums); l > maxSliceLen {
		return fmt.Errorf("flag %s.Enums too long: %d/%d", v.info.DispName, l, maxSliceLen)
	}
	if l := len(v.Ranges); l > 0 {
		if l > maxSliceLen {
			return fmt.Errorf("flag %s.Ranges too long: %d/%d", v.info.DispName, l, maxSliceLen)
		}
		if l%2 != 0 {
			return fmt.Errorf("flag %s.Ranges doesn't match [min,max) pairs: %d", v.info.DispName, l)
		}
		for i := 0; i < l; i += 2 {
			min, max := v.Ranges[i], v.Ranges[i+1]
			if valid := min <= max; !valid {
				return fmt.Errorf("flag %s.Ranges doesn't match [min,max): (%d,%d)", v.info.DispName, min, max)
			}
		}
	}
	if v.Name == "" && v.LogicName == "" {
		return fmt.Errorf("flag missing both Name & LogicName: %v", v)
	}
	if v.Name == "" && len(v.Aliases) > 0 {

	}

	if err := v.validateValues(v.Default); err != nil {
		return fmt.Errorf("default value invalid: %s", err.Error())
	}
	if v.Target != nil {
		v.target = v.Target
	} else {
		v.target = new(float64)
	}
	return nil
}

// IsSet check if value was set
func (v *Float64Flag) IsSet() bool {
	return v.info.HasBeenSet
}

// Apply coordinate the value to flagset
func (v *Float64Flag) Apply(set *flag.FlagSet) error {
	return nil
}

// String return the value for view
func (v *Float64Flag) String() string {
	return ""
}

// ValidateValues verify if all values was valid
func (v *Float64Flag) ValidateValues() error {
	return v.validateValues(*v.target)
}

// Info returns parsed info of this flag
func (v *Float64Flag) Info() *FlagInfo {
	return &v.info
}

// Reset clean the last parsed value of this flag
func (v *Float64Flag) Reset() {
	var t float64
	*v.target = t
	v.info.HasBeenSet = false
}

// for default value verify
func (v *Float64Flag) validateValues(values float64) error {
	return v.validValue(values)
}

// check if value if valid for this flag
func (v *Float64Flag) validValue(value float64) error {
	f := &v.info
	if len(v.Enums) > 0 {
		found := false
		for _, v := range v.Enums {
			if value == v {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("flag %s value %v out of Enums: %v", f.LogicName, value, v.Enums)
		}
	}
	if len(v.Ranges) > 0 {
		found := false
		for i := 0; i < len(v.Ranges); i += 2 {
			min, max := v.Ranges[i], v.Ranges[i+1]
			if value >= min && value < max {
				found = true
				break
			}
		}
		if !found {
			return fmt.Errorf("flag %s value %v out of Ranges: %v", f.LogicName, value, v.Enums)
		}
	}
	return nil
}

var _ Flag = (*Float64Flag)(nil) //for interface verification only
