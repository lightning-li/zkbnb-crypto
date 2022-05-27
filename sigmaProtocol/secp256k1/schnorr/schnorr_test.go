/*
 * Copyright © 2021 Zecrey Protocol
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package schnorr

import (
	"gotest.tools/assert"
	"testing"
	"github.com/zecrey-labs/zecrey-crypto/ecc/zp256"
	"github.com/zecrey-labs/zecrey-crypto/elgamal/secp256k1/twistedElgamal"
)

// pk = g^{sk}
func TestProveVerify(t *testing.T) {
	sk, pk := twistedElgamal.GenKeyPair()
	g := zp256.Base()
	z, A := Prove(sk, g, pk)
	res := Verify(z, A, pk, g)
	assert.Equal(t, true, res)
}