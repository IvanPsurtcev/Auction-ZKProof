import {ethers, network} from "hardhat"
import {Signer, Contract, ContractFactory, BigNumber} from "ethers"
import chai from "chai"
import type {SignerWithAddress} from "@nomiclabs/hardhat-ethers/signers"
import {
    Auction,
    Auction__factory,
} from "../typechain-types";


describe('Auction Contract', () => {
  let provider: MockProvider;
  let auction: Auction;
  let accounts: ethers.Signer[];

  beforeEach(async () => {
    provider = new MockProvider();
    accounts = provider.getWallets();

    auction = await deployContract(1, 1, 1, 1, 1,  accounts[0], auction);
  });

  it('should allow a bidder to make a bid', async () => {
    const bidder = accounts[1];
    const signalForBid = ethers.utils.parseEther('1.0');

    const uniqueHash = 100001;
    await auction.connect(bidder).makeBid(uniqueHash, { value: signalForBid });

    const bidAmounts = await auction.bidAmounts();
    expect(bidAmounts.length).to.equal(1);
  });

  it('should not allow the auctioneer to make a bid', async () => {
    const auctioneer = accounts[0];
    const signalForBid = ethers.utils.parseEther('1.0');

    await expect(
      auction.connect(auctioneer).makeBid(/, { value: signalForBid })
    ).to.be.revertedWith('Auctioneer cant bid');
  });

  it('should not allow a bid without the required signal', async () => {
    const bidder = accounts[1];
    const invalidSignalForBid = ethers.utils.parseEther('0.5');

    const uniqueHash = 100001;
    await expect(
      auction.connect(bidder).makeBid(uniqueHash, { value: invalidSignalForBid })
    ).to.be.revertedWith('You have to pay signal for bidding');
  });

  it('should not allow more bids than the number of bidders', async () => {
    const bidder = accounts[1];
    const signalForBid = ethers.utils.parseEther('1.0');

    for (let i = 0; i < numberOfBidders; i++) {
      await auction.connect(bidder).makeBid( { value: signalForBid });
    }

    await expect(
      auction.connect(bidder).makeBid(commitment, { value: signalForBid })
    ).to.be.revertedWith('All the possible bids has been done, no more bids can be done');
  });
});
