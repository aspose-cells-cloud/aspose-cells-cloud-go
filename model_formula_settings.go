/** --------------------------------------------------------------------------------------------------------------------
* <copyright company="Aspose" file="FormulaSettings.go">
*   Copyright (c) 2025 Aspose.Cells Cloud
* </copyright>
* <summary>
*   Permission is hereby granted, free of charge, to any person obtaining a copy
*  of this software and associated documentation files (the "Software"), to deal
*  in the Software without restriction, including without limitation the rights
*  to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
*  copies of the Software, and to permit persons to whom the Software is
*  furnished to do so, subject to the following conditions:
* 
*  The above copyright notice and this permission notice shall be included in all
*  copies or substantial portions of the Software.
* 
*  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
*  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
*  FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
*  AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
*  LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
*  OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
*  SOFTWARE.
* </summary> 
-------------------------------------------------------------------------------------------------------------------- **/

package asposecellscloud

type FormulaSettings struct {
 
    CalculateOnOpen bool `json:"CalculateOnOpen,omitempty" xml:"CalculateOnOpen"`
    CalculateOnSave bool `json:"CalculateOnSave,omitempty" xml:"CalculateOnSave"`
    ForceFullCalculation bool `json:"ForceFullCalculation,omitempty" xml:"ForceFullCalculation"`
    CalculationMode string `json:"CalculationMode,omitempty" xml:"CalculationMode"`
    CalculationId string `json:"CalculationId,omitempty" xml:"CalculationId"`
    EnableIterativeCalculation bool `json:"EnableIterativeCalculation,omitempty" xml:"EnableIterativeCalculation"`
    MaxIteration int64 `json:"MaxIteration,omitempty" xml:"MaxIteration"`
    MaxChange float64 `json:"MaxChange,omitempty" xml:"MaxChange"`
    PrecisionAsDisplayed bool `json:"PrecisionAsDisplayed,omitempty" xml:"PrecisionAsDisplayed"`
    EnableCalculationChain bool `json:"EnableCalculationChain,omitempty" xml:"EnableCalculationChain"`
    PreservePaddingSpaces bool `json:"PreservePaddingSpaces,omitempty" xml:"PreservePaddingSpaces"`
}
