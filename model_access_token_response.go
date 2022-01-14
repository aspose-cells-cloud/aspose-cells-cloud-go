/*
 *  Copyright (c) 2022 Aspose.Cells Cloud
 *  Permission is hereby granted, free of charge, to any person obtaining a copy
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
 *
 */

package asposecellscloud

type AccessTokenResponse struct {
	Expires string `json:".expires,omitempty" xml:".expires"`
	AccessToken string `json:"access_token,omitempty" xml:"access_token"`
	Issued string `json:".issued,omitempty" xml:".issued"`
	ClientRefreshTokenLifeTimeInMinutes string `json:"clientRefreshTokenLifeTimeInMinutes,omitempty" xml:"clientRefreshTokenLifeTimeInMinutes"`
	ExpiresIn int64 `json:"expires_in,omitempty" xml:"expires_in"`
	TokenType string `json:"token_type,omitempty" xml:"token_type"`
	ClientId string `json:"client_id,omitempty" xml:"client_id"`
	RefreshToken string `json:"refresh_token,omitempty" xml:"refresh_token"`
}
