function subdomainVisits(cpdomains) {
    var visits = {};
    cpdomains.forEach(function (pair) {
        var _a = pair.split(" "), count = _a[0], domain = _a[1];
        var subdomains = domain.split(".");
        var subdomain = "";
        for (var i = subdomains.length - 1; i >= 0; i--) {
            subdomain = subdomains[i] + subdomain;
            if (visits[subdomain] === undefined) {
                visits[subdomain] = parseInt(count);
            }
            else {
                visits[subdomain] += parseInt(count);
            }
            subdomain = "." + subdomain;
        }
    });
    var output = [];
    Object.keys(visits).forEach(function (subdomain) {
        output.push("".concat(visits[subdomain].toString(), " ").concat(subdomain));
    });
    return output;
}
console.log(subdomainVisits([
    "900 google.mail.com",
    "50 yahoo.com",
    "1 intel.mail.com",
    "5 wiki.org",
]));
