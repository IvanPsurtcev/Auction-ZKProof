pragma solidity ^0.8.17;

import "./Verifier.sol";

contract Auction is Verifier {
    Verifier public verifierContractZoKrates;

    uint256 public auctionStartingDate;
    uint256 public auctionEndingDate;

    uint256 public provingBidsStartingDate;
    uint256 public provingBidsEndingDate;
    
    uint256 public signalForBid;
    uint256 public idAuction; 
    uint256 public numberOfBidders; 
    address payable public beneficiary; 
    address public auctioneer; 
    address public winner; 
    bool public auctionEnded; 
    bool public winnerObtained;

    mapping(bytes => address payable) bidsAddresses; 
    bytes32[] public bids; 
    bytes[] public bidAmounts; 
    bytes32[2][] public hashZokratesBids; 
    uint256 public biggestBid; 
    uint256 public positionWinnerBid; 
    bool public expectedHashes;

    constructor(uint256 _auctionStartingDate, uint256 _auctionEndingDate, uint256 _signalForBid,uint256 _idAuction, uint256 _numberOfBidders, address payable _beneficiary, address payable _auctioneer) {
        auctionStartingDate = _auctionStartingDate;
        auctionEndingDate = _auctionEndingDate;
        provingBidsStartingDate = _auctionStartingDate + 200;
        provingBidsEndingDate = _auctionEndingDate + 200;
        signalForBid = _signalForBid * 1e18;
        idAuction = _idAuction;
        beneficiary = _beneficiary;
        auctioneer = _auctioneer;
        numberOfBidders = _numberOfBidders;
    }

    function publishProduct(address payable _verifierContractAddress) external {
        verifierContractZoKrates = Verifier(_verifierContractAddress);
    }

    function makeBid(bytes32 hashedEncryptedBid) public payable{
        require(msg.sender != auctioneer, "Auctioneer cant bid");
        require(msg.value==signalForBid, "You have to pay signal for bidding");
        require(bidAmounts.length < numberOfBidders + 1, "All the possible bids has been done, no more bids can be done");
        bids.push(hashedEncryptedBid);
    }

    function bidProver(bytes32 bidHashedSent, bytes memory encryptedBid, bytes32 hashZokrates1, bytes32 hashZokrates2) public payable{
        require(bidAmounts.length < numberOfBidders + 1, "All the bids have been proved");
        
        for (uint256 i = 0; i<bids.length; i++) {
            if (bids[i] == bidHashedSent) {
                require(bidHashedSent == keccak256Hash(encryptedBid, hashZokrates1, hashZokrates2), "The hashed bid sent does not correspond to hash of the values you are providing");
                bidAmounts.push(encryptedBid);
                hashZokratesBids.push([hashZokrates1, hashZokrates2]);
                bidsAddresses[encryptedBid] = payable(msg.sender);
                delete bids[i];
                break;
            }
        }  
    }
    
    function checkingAuctioneerHashesInputs(uint256 hash0_0, uint256 hash0_1, uint256 hash1_0, uint256 hash1_1, uint256 hash2_0, uint256 hash2_1, uint256 hash3_0, uint256 hash3_1) public payable returns (bool correctHashes){
        uint256[2][4] memory hashes = [[hash0_0, hash0_1], [hash1_0, hash1_1], [hash2_0, hash2_1], [hash3_0, hash3_1]];
        bool sameHashes;
        for(uint256 row = 0; row<hashZokratesBids.length; row++){
            if (!(((keccak256(abi.encodePacked(uint256(hashZokratesBids[row][0])))) == (keccak256(abi.encodePacked((uint256(hashes[row][0])))))) && ((keccak256(abi.encodePacked((uint256(hashZokratesBids[row][1]))))) == (keccak256(abi.encodePacked(uint256(hashes[row][1]))))))){
                return false;
            }
            if(row == 3){
                sameHashes = true;
            }
        }
        expectedHashes = sameHashes;
        return sameHashes;

    }

    function auctionEnd(uint256[2] memory a,
            uint256[2][2] memory b,
            uint256[2] memory c,
            uint256[14] memory input) public payable returns (uint256 position, uint256 highestBid){
        require(msg.sender == auctioneer, "Only the Auctioneer can determine the Winner of the bid");
        require(auctionEnded == false, "The auction has ended");
        bool correctHashesUsedByAuctioneer = checkingAuctioneerHashesInputs(input[0], input[1], input[2], input[3], input[4], input[5], input[6], input[7]);
        require(correctHashesUsedByAuctioneer, "The inputs regarding to the ZoKrates hashes are wrong, put them correctly in order in the array");        bool result = verifierContractZoKrates.verifyTx(a, b, c, input);
        require(result == true, "Incorrect proof");
        biggestBid = input[13] * 1e18;
        positionWinnerBid = input[12];
        winner = bidsAddresses[getBidAmounts(positionWinnerBid)];
        for (uint256 i = 0; i < bidAmounts.length; i++){
            bidsAddresses[getBidAmounts(i)].transfer(signalForBid);
        }
        winnerObtained = true;
        return (input[12], input[13]);
    }

    function paymentOperations() public payable{
        require(auctionEnded == false, "The auction has already ended and the profits are being sent");
        require(winnerObtained == true, "The winner has not been obtained yet");
        require(msg.sender == winner, "You are trying to call this function but you are not the winner");
        require(msg.value == biggestBid, "You have to pay the bid you promised for obtaining your profits");
        beneficiary.transfer(msg.value);
        auctionEnded = true;

    }

    function getBidAmounts(uint256 i) public view returns(bytes memory){
        return bidAmounts[i];
    }

    function keccak256Hash(bytes memory encrypted, bytes32 hashZokrates1, bytes32 hashZokrates2) public pure returns (bytes32 hashSolidity){
        return keccak256(abi.encode(encrypted, hashZokrates1, hashZokrates2));
    }

    function setExpectedHashes(bool isCorrect) public {
        expectedHashes = isCorrect;
    }
}