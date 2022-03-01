class HashTable {
  constructor(size) {
    this.data = new Array(size);
  }

  _hash(key) {
    let hash = 0;
    for (let i = 0; i < key.length; i++) {
      hash = (hash + key.charCodeAt(i)) % this.data.length;
    }
    return hash;
  }

  set(key, val) {
    const address = this._hash(key);
    if (!this.data[address]) {
      this.data[address] = [[key, val]];
    } else {
      this.data[address].push([key, val]);
    }
  }

  get(key) {
    const address = this._hash(key);
    const currentBucket = this.data[address];
    for (let i = 0; i < currentBucket.length; i++) {
      if (currentBucket[i][0] === key) {
        return currentBucket[i];
      }
    }
    return undefined;
  }

  keys() {
    const keysArray = [];
    for (let i = 0; i < this.data.length; i++) {
      const currentBucket = this.data[i];
      if (currentBucket) {
        for (let j = 0; j < currentBucket.length; j++) {
          keysArray.push(currentBucket[j][0]);
        }
      }
    }
    return keysArray;
  }
}

const myHashTable = new HashTable(2);
myHashTable.set("grape0", 1);
myHashTable.set("grape1", 2);
myHashTable.set("grape2", 3);
console.log(myHashTable.get("grape20"));
console.log(myHashTable.keys());
