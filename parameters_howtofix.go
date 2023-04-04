// Copyright 2023 Princess B33f Heavy Industries / Dave Shanley
// SPDX-License-Identifier: MIT

package main

const (
    HowToFixReservedValues string = "parameter values need to URL Encoded to ensure reserved " +
        "values are correctly encoded, for example: '%s'"
    HowToFixParamInvalidNumber  string = "Convert the value '%s' into a number"
    HowToFixParamInvalidBoolean string = "Convert the value '%s' into a true/false value"
    HowToFixParamInvalidExplode string = "When setting explode to true, make sure each parameter name" +
        " is unique, and the values are separated by commas."
    HowToFixParamInvalidFormEncode                  string = "Use a form style encoding for parameter values, for example: '%s'"
    HowToFixParamInvalidSchema                      string = "Ensure that the object being submitted, matches the schema correctly"
    HowToFixParamInvalidSpaceDelimitedObject        string = "Objects cannot be space delimited. Use a form style encoding, for example: '%s'"
    HowToFixParamInvalidPipeDelimitedObject         string = "Objects cannot be pipe delimited. Use a form style encoding, for example: '%s'"
    HowToFixParamInvalidSpaceDelimitedObjectExplode string = "When using 'explode' with space delimited parameters, " +
        "they should be separated by spaces. For example: '%s'"
    HowToFixParamInvalidPipeDelimitedObjectExplode string = "When using 'explode' with pipe delimited parameters, " +
        "they should be separated by pipes '|'. For example: '%s'"
    HowToFixParamInvalidDeepObjectMultipleValues string = "There can only be a single value per property name, " +
        "deepObject parameters should contain the property key in square brackets next to the parameter name. For example: '%s'"
    HowToFixParamInvalidMatrixMultipleValues string = "There can only be a single value per property name, " +
        "matrix parameters with no 'explode' should use commas to separate values. For example: '%s'"
)
