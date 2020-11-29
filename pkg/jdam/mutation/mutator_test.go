package mutation_test

import (
	"testing"

	"gitlab.com/michenriksen/jdam/pkg/jdam/mutation"
)

func TestMutatorsIncludeRandomInt(t *testing.T) {
	m := &mutation.RandomInt{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include RandomInt")
}

func TestMutatorsIncludeRandomFloat(t *testing.T) {
	m := &mutation.RandomFloat{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include RandomFloat")
}

func TestMutatorsIncludeNumericToNegative(t *testing.T) {
	m := &mutation.NumericToNegative{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include NumericToNegative")
}

func TestMutatorsIncludeIntegerOverflow(t *testing.T) {
	m := &mutation.IntegerOverflow{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include IntegerOverflow")
}

func TestMutatorsIncludeByteSwap(t *testing.T) {
	m := &mutation.ByteSwap{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include ByteSwap")
}

func TestMutatorsIncludeByteInsert(t *testing.T) {
	m := &mutation.ByteInsert{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include ByteInsert")
}

func TestMutatorsIncludeByteDrop(t *testing.T) {
	m := &mutation.ByteDrop{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include ByteDrop")
}

func TestMutatorsIncludeByteRepeat(t *testing.T) {
	m := &mutation.ByteRepeat{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include ByteRepeat")
}

func TestMutatorsIncludeByteSequenceRepeat(t *testing.T) {
	m := &mutation.ByteSequenceRepeat{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include ByteSequenceRepeat")
}

func TestMutatorsIncludeBoolInvert(t *testing.T) {
	m := &mutation.BoolInvert{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include BoolInvert")
}

func TestMutatorsIncludeSQLInjection(t *testing.T) {
	m := &mutation.SQLInjection{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include SQLInjection")
}

func TestMutatorsIncludeCommandInjection(t *testing.T) {
	m := &mutation.CommandInjection{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include CommandInjection")
}

func TestMutatorsIncludeXPATHInjection(t *testing.T) {
	m := &mutation.XPATHInjection{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include XPATHInjection")
}

func TestMutatorsIncludeLDAPInjection(t *testing.T) {
	m := &mutation.LDAPInjection{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include LDAPInjection")
}

func TestMutatorsIncludeNoSQLInjection(t *testing.T) {
	m := &mutation.NoSQLInjection{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include NoSQLInjection")
}

func TestMutatorsIncludeFormatStringInjection(t *testing.T) {
	m := &mutation.FormatStringInjection{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include FormatStringInjection")
}

func TestMutatorsIncludeLocalFileInclusion(t *testing.T) {
	m := &mutation.LocalFileInclusion{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include LocalFileInclusion")
}

func TestMutatorsIncludeTemplateInjection(t *testing.T) {
	m := &mutation.TemplateInjection{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include TemplateInjection")
}

func TestMutatorsIncludeSillyString(t *testing.T) {
	m := &mutation.SillyString{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include SillyString")
}

func TestMutatorsIncludeSillyObject(t *testing.T) {
	m := &mutation.SillyObject{}
	for _, v := range mutation.Mutators {
		if v.ID() == m.ID() {
			return
		}
	}
	t.Error("Expected Mutators to include SillyObject")
}
