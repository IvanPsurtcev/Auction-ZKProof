import { ethers } from "hardhat";
import { expect } from "chai";
import { Verifier__factory, Verifier } from "../typechain-types";

describe("Verifier contract", function () {
  let pairing: Verifier;

  beforeEach(async function () {
    const factory = (await ethers.getContractFactory(
      "Verifier"
    )) as Verifier__factory;
    pairing = await factory.deploy();
  });

  it("should add two points correctly", async function () {
    const p1 = { X: 3, Y: 4 };
    const p2 = { X: 2, Y: 5 };
    const expected = { X: 5, Y: 0 };

    const result = await pairing.addition(p1, p2);

    expect(result).to.deep.equal(expected);
  });

  it("should negate a point correctly", async function () {
    const p = { X: 3, Y: 4 };
    const expected = {
      X: 3,
      Y: 21888242871839275222246405745257275088696311157297823662689037894645226208579,
    };

    const result = await pairing.negate(p);

    expect(result).to.deep.equal(expected);
  });

  it("should scalar multiply a point correctly", async function () {
    const p = { X: 3, Y: 4 };
    const s = 5;
    const expected = {
      X: 21817298385435845187866423467870356413319732964554362201200803416706794648928,
      Y: 6476209829095821334707548436506881549702001787173875787129401751863632810669,
    };

    const result = await pairing.scalar_mul(p, s);

    expect(result).to.deep.equal(expected);
  });

  it("should check pairing correctly", async function () {
    const p1 = [
      { X: 1, Y: 2 },
      { X: 3, Y: 4 },
    ];
    const p2 = [
      {
        X: [
          11559732032986387107991004021392285783925812861821192530917403151452391805634,
          10857046999023057135944570762232829481370756359578518086990519993285655852781,
        ],
        Y: [
          4082367875863433681332203403145435568316851327593401208105741076214120093531,
          8495653923123431417604973247489272438418190587263600148770280649306958101930,
        ],
      },
      {
        X: [
          12623981093529297811592795440556916332152627681365585961155223559685317870299,
          18454937603337218674455881075897601955491279027326184678846403196817384281992,
        ],
        Y: [
          1917725148804502947097955756941785610190543234038743038260825796810475764530,
          1933937151057647945729335973036523017066367221787844784541668068447736450006,
        ],
      },
    ];
    const expected = true;

    const result = await pairing.pairing(p1, p2);

    expect(result).to.equal(expected);
  });

  it("Correct mint and all functions", async function () {
    const { aayc, owner, user1, user2, user3 } = await loadFixture(deploy);

    const whitelistFirstAddresses = [owner.address, user1.address];
    const merkleTreeFirst = generateTree(whitelistFirstAddresses);

    let whitelistSecondAddresses = [
      owner.address,
      user1.address,
      user2.address,
    ];
    const merkleTreeSecond = generateTree(whitelistSecondAddresses);

    expect(
      await aayc.connect(owner).setMerkleRootFirst(merkleTreeFirst.getHexRoot())
    ).to.be.ok;
    expect(
      await aayc
        .connect(owner)
        .setMerkleRootSecond(merkleTreeSecond.getHexRoot())
    ).to.be.ok;

    const hexProofOwnerFirst = getProof(merkleTreeFirst, owner.address);
    const hexProofUser1First = getProof(merkleTreeFirst, user1.address);
    const hexProofUser2First = getProof(merkleTreeFirst, user2.address);
    const hexProofUser3First = getProof(merkleTreeFirst, user3.address);

    const hexProofOwnerSecond = getProof(merkleTreeSecond, owner.address);
    const hexProofUser1Second = getProof(merkleTreeSecond, user1.address);
    const hexProofUser2Second = getProof(merkleTreeSecond, user2.address);
    const hexProofUser3Second = getProof(merkleTreeSecond, user3.address);

    const provider = ethers.provider;
    expect(await provider.getBalance(aayc.address)).to.equal(0);

    await expect(
      aayc
        .connect(user1)
        .mintWhitelistFirst(2, hexProofUser1First, { value: ether("0.1") })
    ).to.be.revertedWith("Auction is closed");
    await expect(
      aayc
        .connect(user1)
        .mintWhitelistSecond(2, hexProofUser1First, { value: ether("0.1") })
    ).to.be.revertedWith("Auction is closed");
});
