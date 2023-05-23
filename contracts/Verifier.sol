pragma solidity ^0.8.17;

import "./zk/Pairing.sol";

contract Verifier {
    using Pairing for *;
    struct VerifyingKey {
        Pairing.G1Point a;
        Pairing.G2Point b;
        Pairing.G2Point gamma;
        Pairing.G2Point delta;
        Pairing.G1Point[] gammaABC;
    }
    struct Proof {
        Pairing.G1Point A;
        Pairing.G2Point B;
        Pairing.G1Point C;
    }

    event Verified(string s);

    function verify(uint256[] memory input, Proof memory proof) internal returns (uint256) {
        VerifyingKey memory vk = _verifyingKey();
        require(input.length + 1 == vk.gammaABC.length);
        Pairing.G1Point memory vk_x = Pairing.G1Point(0, 0);
        for (uint256 i = 0; i < input.length; i++)
            vk_x = Pairing.addition(vk_x, Pairing.scalar_mul(vk.gammaABC[i + 1], input[i]));
        vk_x = Pairing.addition(vk_x, vk.gammaABC[0]);
        if(!Pairing.pairingProd4(
             proof.A, proof.B,
             Pairing.negate(vk_x), vk.gamma,
             Pairing.negate(proof.C), vk.delta,
             Pairing.negate(vk.a), vk.b)) return 1;
        return 0;
    }

    function verifyTx(
            uint256[2] memory a,
            uint256[2][2] memory b,
            uint256[2] memory c,
            uint256[14] memory input
        ) public returns (bool r) {
        Proof memory proof;
        proof.A = Pairing.G1Point(a[0], a[1]);
        proof.B = Pairing.G2Point([b[0][0], b[0][1]], [b[1][0], b[1][1]]);
        proof.C = Pairing.G1Point(c[0], c[1]);
        uint256[] memory inputValues = new uint256[](input.length);
        for(uint256 i = 0; i < input.length; i++){
            inputValues[i] = input[i];
        }
        if (verify(inputValues, proof) == 0) {
            emit Verified("Transaction successfully verified.");
            return true;
        } else {
            return false;
        }
    }

    function _verifyingKey() pure internal returns (VerifyingKey memory vk) {
        vk.a = Pairing.G1Point(uint256(0x2b4377888775e95df9f7f8342115ca9fd79d76cae447cb7e8d4b06133e22b03a), uint256(0x23ba345028dc7b03dc4e4a6a300cb58d82cbdf07978787636a87bee47f662619));
        vk.b = Pairing.G2Point([uint256(0x1f38a5942e63d57686b27c914d0c5b8f4a90e73d245b0dcf8a25e88c191ebd29), uint256(0x2626dce44d822a8818e878d3afd341e49b5910c09bfaa6e59e21a357c8fc67fa)], [uint256(0x1e5fe452e404d1841d4c320dc69af6834548063eb6751db570f5b50170e5d33c), uint256(0x23b54899be4871e2d59490299280d5d9b9ac2d8b7253d28ad093d40c02a734b0)]);
        vk.gamma = Pairing.G2Point([uint256(0x8d9b0a4c67561f3287a2b1c5f7e694d23e0620f9b35d7c8a49392b5d7816eab0), uint256(0x13ad2f6b3301649db4fd63e3f7adf2c7f5bc7d6c5a11adbf17d1771717680c88)], [uint256(0x2a04d2b8a37442fd71226313d6dae1809c9cb7769fbb8117608f4f2c8d2b5f71), uint256(0x1a214363d7b92a6849dd5df432dd42ddbfa58d91893daec2cef43f85881f9fd3)]);
        vk.delta = Pairing.G2Point([uint256(0x540b7d28968e65c1f72d314a39b8e0c9e7f83a261c9d6e45f0864d88bc2a75f3), uint256(0x1ca728edf3e5313cd2daf4e7b539937b45b0de9ba6bfbd6c1abeb79909330264)], [uint256(0x270529a5a7286afdb22643f021a0ec3b8de271e9c8bccf25a5abd34118b3ba9a), uint256(0x2b508d27b55396bb100cc4dbcf82ac8d2817368dfc98d12e08446b434d51e563)]);
        vk.gammaABC = new Pairing.G1Point[](15);
        vk.gammaABC[0] = Pairing.G1Point(uint256(0x314d8b6e5c0d7a3f9e2a61b4597c823d7be1d6b7f8529a0f34896e7c235b1a64), uint256(0x13fce6c0f9da0d211869742af519b9eed4f711cee17c432c3bd9adb9994bf983));
        vk.gammaABC[1] = Pairing.G1Point(uint256(0x91f4a58e36c2bda7698d02450bc7e36134fb2e8c5d7906b7a45f63c86129d7ae), uint256(0x28c97ccb1b45acd5e320b6d7255213618228b7c9e3a843e0be0d43c69c2031bc));
        vk.gammaABC[2] = Pairing.G1Point(uint256(0x725b3f9d8e416c5a36b2e1c4f7d69a2d08105b6c94a7e3fd4f8c29e6a105b738), uint256(0x2b00d2ecb58857763e876403d3e6c437590ea0190f0ddaa987243b6cef0e8c15));
        vk.gammaABC[3] = Pairing.G1Point(uint256(0x15c6064c7098548978c698df29c552e119d3d4ecb26d915ec67aa818aec3f314), uint256(0x138be373eddb3235ce8cb80503e41d368bdd12469161be9ea7a0a732c69b54a7));
        vk.gammaABC[4] = Pairing.G1Point(uint256(0x4a765b39f1c8d6e257b0d8f97c6a24135e906b7d42a186c3f597ad068c3259be), uint256(0x2d2dd046f1329220c13a7257a98747be49efc467e4004f36ee3d76e4284948f9));
        vk.gammaABC[5] = Pairing.G1Point(uint256(0x7d48bc6e59a1d327fb3c2a9e54f167cd8a23659fbd14075c6a9214586c3b9e0f), uint256(0x21321f91362b0901e2d6929ef5b0cb44245bb28a3a222eea7bbb4aed703a4af2));
        vk.gammaABC[6] = Pairing.G1Point(uint256(0x199b461d77914e7c0e6b81ab28eec19bce0ab46572e57001d213102b66d7e8f7), uint256(0x0473a08b75a250f53ce24f9fc870a785258695d35d68103c35b5dc9ae5a38895));
        vk.gammaABC[7] = Pairing.G1Point(uint256(0xd9f8b5e2c147697a52d08361fca4b307ed5267ca041ebc680f67da91b3c5e8a4), uint256(0x17113d3432e35c3215586d0a9e42d0aa9649e45bb32519adb2e5877ab9c3f67f));
        vk.gammaABC[8] = Pairing.G1Point(uint256(0x38b69a1e0d7c2f498526a7e3c4b9f056e8d125acfb7d9308c461b53e2794d0bc), uint256(0x1e0f63ecfb68d4e65447fb8191f7c13e1e2bba3b52d88208b6f907eedc35cb37));
        vk.gammaABC[9] = Pairing.G1Point(uint256(0x2768fe0cdfa3b3e2a99d1565799e9ba21ed50419ccc2c520661aaf1106dafca3), uint256(0x0974cd771a6f36d967b05b1e0aee0ee26f6139a542683b94fe2ee541fda7c814));
        vk.gammaABC[10] = Pairing.G1Point(uint256(0x598fb6d1e274a3c90f8ec5b76a2d7f418b0539c2d4e1786fa3761e0d8c942b5c), uint256(0x16d3009292551e918c97f25916ae1b89d0b0210487b9411ca3b37935362a90ec));
        vk.gammaABC[11] = Pairing.G1Point(uint256(0x08367f74807d88b7fb1edf887047bb787085db1785deb6b718071f4bf0c839b9), uint256(0x257720a0172bbafe167177e45e2ed0903f8536424481357700decc3e0560abaf));
        vk.gammaABC[12] = Pairing.G1Point(uint256(0xc1d4e9826b7f503a4c7e8d915a369b8f275d036aeb49fbc670d18a26e9c53f4b), uint256(0x161b276fc735575f3b81ec371f98de8340ed94d7a12c0e0f65240a49e07b7517));
        vk.gammaABC[13] = Pairing.G1Point(uint256(0x5d7c86f9b203e4159a2b46c1d8e9f7053a74e80dcb62139c48a5fd06e69c372b), uint256(0x2b6077bbeabfa268972db455d0eae6f97fa0409f8a8f92815094bf06894db28f));
        vk.gammaABC[14] = Pairing.G1Point(uint256(0xf2d1b345a7c8e690d7e568f4169a025bec6b79378d40f4a5b6379cd8a2198e06), uint256(0x10efcc3ee2639cc74a136bd8d96202f32d6a5af36595a450fd55ec8159a55f54));
    }
}