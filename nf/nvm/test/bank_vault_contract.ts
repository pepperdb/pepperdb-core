class DepositeContent {
    balance: BigNumber;
    expiryHeight: BigNumber;

    constructor(text?: string) {
        if (text) {
            let o = JSON.parse(text);
            this.balance = new BigNumber(o.balance);
            this.expiryHeight = new BigNumber(o.expiryHeight);
        } else {
            this.balance = new BigNumber(0);
            this.expiryHeight = new BigNumber(0);
        }
    }

    toString(): string {
        return JSON.stringify(this);
    }

}

class BankVaultContract {
    constructor() {
        LocalContractStorage.defineMapProperty(this, "bankVault", {
            parse(text: string): DepositeContent {
                return new DepositeContent(text);
            },

            stringify(o: DepositeContent): string {
                return o.toString();
            },
        });
    }

    // init.
    init() {
        // pass.
    }

    // save.
    save(height: number) {
        let from = Blockchain.transaction.from;
        let value = Blockchain.transaction.value;
        let bk_height = new BigNumber(Blockchain.block.height);

        let orig_deposit = this.bankVault.get(from);

        if (orig_deposit) {
            value = value.plus(orig_deposit.balance);
        }

        let deposit = new DepositeContent();
        deposit.balance = value;
        deposit.expiryHeight = bk_height.plus(height);

        this.bankVault.put(from, deposit);
    }

    takeout(value: number) {
        let from = Blockchain.transaction.from;
        let bk_height = new BigNumber(Blockchain.block.height);
        let amount = new BigNumber(value);

        let deposit = this.bankVault.get(from);
        if (!deposit) {
            throw new Error("No deposit before.");
        }

        if (bk_height.lt(deposit.expiryHeight)) {
            throw new Error("Can't takeout before expiryHeight.");
        }

        if (amount.gt(deposit.balance)) {
            throw new Error("Insufficient balance.");
        }

        let result = Blockchain.transfer(from, amount);
        if (result == false) {
            throw new Error("transfer failed.");
        }

        deposit.balance = deposit.balance.sub(amount);
        this.bankVault.put(from, deposit);
    }

    balanceOf() {
        let from = Blockchain.transaction.from;
        return this.bankVault.get(from);
    }
}

module.exports = BankVaultContract;
