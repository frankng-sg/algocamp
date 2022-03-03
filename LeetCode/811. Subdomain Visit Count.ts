function subdomainVisits(cpdomains: string[]): string[] {
  let visits = {};

  cpdomains.forEach((pair) => {
    const [count, domain] = pair.split(" ");
    const subdomains = domain.split(".");
    let subdomain = "";
    for (let i = subdomains.length - 1; i >= 0; i--) {
      subdomain = subdomains[i] + subdomain;
      if (visits[subdomain] === undefined) {
        visits[subdomain] = parseInt(count);
      } else {
        visits[subdomain] += parseInt(count);
      }
      subdomain = "." + subdomain;
    }
  });

  const output = [];
  Object.keys(visits).forEach((subdomain) => {
    output.push(`${visits[subdomain].toString()} ${subdomain}`);
  });

  return output;
}

console.log(
  subdomainVisits([
    "900 google.mail.com",
    "50 yahoo.com",
    "1 intel.mail.com",
    "5 wiki.org",
  ])
);
