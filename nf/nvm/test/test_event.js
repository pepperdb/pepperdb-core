Event.Trigger("ERC20", {
    Transfer: {
        from: "0x0",
        to: "0x1",
        value: 1234,
    }
});

Event.Trigger("ERC20", {
    Issue: {
        to: "0x1",
        value: 2234,
    }
});
