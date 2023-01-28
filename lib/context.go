package lib

import (
	"errors"
	"fmt"
	"gchain/help"
	"strings"
)

type Context struct {
	frozen bool
	param  values
	data   values
}

type values map[string]interface{}

func NewContext() Context {
	return Context{
		param: make(values),
		data:  make(values),
	}
}

func (context *Context) AddSingleParam(paramKey string, paramValue interface{}) error {
	if context.frozen {
		return errors.New("the context is frozen")
	}
	if len(paramKey) == 0 || paramValue == nil {
		return errors.New("pls enter the param")
	}
	context.param[paramKey] = paramValue
	return nil
}

func (context *Context) AddMultiParam(param map[string]interface{}) error {
	if context.frozen {
		return errors.New("the context is frozen")
	}
	if param == nil || len(param) == 0 {
		return errors.New("pls enter the param")
	}
	for paramKey, paramValue := range param {
		err := context.AddSingleParam(paramKey, paramValue)
		if err != nil {
			return err
		}
	}
	return nil
}

func (context *Context) AddSingleData(dataKey string, dataValue interface{}) error {
	if len(dataKey) == 0 || dataValue == nil {
		return errors.New("pls enter the data")
	}
	context.data[dataKey] = dataValue
	return nil
}

func (context *Context) AddMultiData(data map[string]interface{}) error {
	if data == nil || len(data) == 0 {
		return errors.New("pls enter the param")
	}
	for dataKey, dataValue := range data {
		err := context.AddSingleData(dataKey, dataValue)
		if err != nil {
			return err
		}
	}
	return nil
}

const SUMMARY_VALUES_LIMIT = 5
const SUMMARY_PRETTY_FMT = `
the context is %s.
the param contains (limit 5): 
%s
the data contains (limit 5):
%s
`

func (context Context) Summary() string {
	return fmt.Sprintf(
		SUMMARY_PRETTY_FMT,
		help.If(context.frozen, "frozen", "not frozen"),
		summaryValue(context.param),
		summaryValue(context.data),
	)
}

func summaryValue(value values) string {
	oversize := SUMMARY_VALUES_LIMIT < len(value)
	valueLimit := help.If(oversize, SUMMARY_VALUES_LIMIT, len(value)).(int)
	var valueSummary strings.Builder
	for k, v := range value {
		valueSummary.WriteString("  ")
		valueSummary.WriteString(k)
		valueSummary.WriteString(" : ")
		valueSummary.WriteString(fmt.Sprint(v))
		valueLimit--
		if valueLimit == 0 {
			if oversize {
				valueSummary.WriteString("\n")
				valueSummary.WriteString("  ")
				valueSummary.WriteString("...")
			}
			break
		}
		valueSummary.WriteString("\n")
	}
	return valueSummary.String()
}
