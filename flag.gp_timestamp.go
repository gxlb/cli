///////////////////////////////////////////////////////////////////
//
// !!!!!!!!!!!! NEVER MODIFY THIS FILE MANUALLY !!!!!!!!!!!!
//
// This file was auto-generated by tool [github.com/gxlb/gogp]
// Last update at: [Sat Mar 20 2021 22:32 CST]
// Generate from:
//   [github.com/gxlb/cli/internal/gp/flag.gp]
//   [github.com/gxlb/cli/flag.gpg] [flag_timestamp]
//
// Tool [github.com/gxlb/gogp] info:
// CopyRight 2021 @Ally Dale. All rights reserved.
// Author  : Ally Dale(vipally@gmail.com)
// Site    : https://github.com/vipally
// Version : v4.0.0
// 
///////////////////////////////////////////////////////////////////

// MIT License
//
// Copyright (c) 2021 Ally Dale <vipally@gamil.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package cli

import (
	"flag"
	"fmt"

	"cli/internal/impl"
	"cli/internal/util"
)

// TimestampFlag define a value of type *impl.TimestampValue
type TimestampFlag struct {
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
	Target      *impl.TimestampValue // Target value pointer outside
	Default     *impl.TimestampValue // Default value
	DefaultText string               // Default value display in help info
	Layout      string               // Layout of the time format, default **2006-01-02T15:04:05 MST**
	//
	////////////////////////////////////////////////////////////////////////////
	//area for parsing
	target *impl.TimestampValue // target value pointer(maybe new(*impl.TimestampValue) if Target not set)
	info   impl.FlagInfo        // parsed info of this flag
}

// init verify and init the flag info
func (f *TimestampFlag) init(namegen *util.NameGenenerator) error {
	f.info.Flag = f
	f.info.EnvVars = append([]string{}, f.EnvVars...)
	f.info.Usage = f.Usage
	f.info.DefaultText = f.DefaultText
	f.info.Required = f.Required
	f.info.Hidden = f.Hidden
	f.info.FilePath = f.FilePath
	f.info.HasBeenSet = false
	f.info.Name = namegen.GetOrGenName(f.Name)
	f.info.NonameFlag = f.info.Name != f.Name
	f.info.LogicName = impl.FlagLogicName(f.Name, f.LogicName)
	f.info.ValueName = impl.FlagValueName(f.LogicName)
	impl.MergeNames(f.info.Name, f.Aliases, &f.info.Names) //use v.info.Name to enable auto-generated name

	//make the target pointer
	if f.Target != nil {
		f.target = f.Target
	} else {
		f.target = impl.NewEmptyTimestampValue(f.Layout)
	}

	if f.Name == "" && f.LogicName == "" { // Name & LogicName cannot both missing
		return fmt.Errorf("flag missing both Name & LogicName: %v", f)
	}
	if f.Name == "" && len(f.Aliases) > 0 { // Noname ones must without Aliases
		return fmt.Errorf("flag %s missing name, but has Aliases %v", f.info.LogicName, f.Aliases)
	}
	if err := f.validateValues(f.Default); err != nil { // verify default values
		return fmt.Errorf("default value invalid: %s", err.Error())
	}
	if err := util.FiltNames(f.info.Names); err != nil { // verify name duplicate
		return fmt.Errorf("flag %s.Names error: %s", f.info.LogicName, err.Error())
	}
	if err := util.FiltNames(f.info.EnvVars); err != nil { // verify EnvVars duplicate
		return fmt.Errorf("flag %s.EnvVars error: %s", f.info.LogicName, err.Error())
	}
	return nil
}

// IsSet check if value was set
func (f *TimestampFlag) IsSet() bool {
	return f.info.HasBeenSet
}

//GetLogicName returns the logic name of the falg
func (f *TimestampFlag) GetLogicName() string {
	return f.info.LogicName
}

//GetValueName returns the value name of the falg
func (f *TimestampFlag) GetValueName() string {
	return f.info.ValueName
}

// Names returns the names of the flag
func (f *TimestampFlag) Names() []string {
	return f.info.Names
}

// IsRequired returns whether or not the flag is required
func (f *TimestampFlag) IsRequired() bool {
	return f.Required
}

// TakesValue returns true of the flag takes a value, otherwise false
func (f *TimestampFlag) TakesValue() bool {
	return false
}

// GetUsage returns the usage string for the flag
func (f *TimestampFlag) GetUsage() string {
	return f.info.Usage
}

// GetValue returns the flags value as string representation.
func (f *TimestampFlag) GetValue() string {
	return ""
}

// Apply coordinate the value to flagset
func (f *TimestampFlag) Apply(set *flag.FlagSet) error {
	return nil
}

// String return the value for view
func (f *TimestampFlag) String() string {
	return ""
}

// ValidateValues verify if all values are valid
func (f *TimestampFlag) ValidateValues() error {
	return f.validateValues(f.target)
}

// Info returns parsed info of this flag, the returned object must READ-ONLY
func (v *TimestampFlag) Info() *impl.FlagInfo {
	return &v.info
}

// Reset clean the last parsed value of this flag
func (f *TimestampFlag) Reset() {
	f.target.Reset()
	f.info.HasBeenSet = false
}

// for default value verify
func (f *TimestampFlag) validateValues(values *impl.TimestampValue) error {
	return f.validValue(values)
}

// check if value if valid for this flag
func (f *TimestampFlag) validValue(value *impl.TimestampValue) error {
	return nil
}

// Timestamp looks up the value of a local TimestampFlag
func (c *Context) Timestamp(name string) *impl.TimestampValue {
	if fs := lookupFlagSet(name, c); fs != nil {
		return lookupTimestamp(name, fs)
	}
	return nil
}

func lookupTimestamp(name string, set *flag.FlagSet) *impl.TimestampValue {
	f := set.Lookup(name)
	if f != nil {
		//TODO:
	}
	return nil
}

var _ impl.Flag = (*TimestampFlag)(nil) //for interface verification only
