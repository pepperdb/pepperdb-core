declare var require: {
    (id: string): any;
}

declare var Console: {
    prototype: Console;
    Console: Console;
};

declare var ContractStorage: {
    prototype: ContractStorage;
    lcs: ContractStorage;
    gcs: ContractStorage;
}
declare var LocalContractStorage: ContractStorage;
declare var GlobalContractStorage: ContractStorage;
