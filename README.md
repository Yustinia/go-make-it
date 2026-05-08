# go-make-it

A CLI command written in GO to automatically create a standardized project folder structure

---

## Build

**Using `just`:**
```bash
just build
```

**Using `go build`:**
```bash
go build -o gmk .
```

Both produce a binary named `gmk` in the current directory.

---

## Install

**Using `just`:**
```bash
just install
```

Automatically builds and moves the binary to `~/.local/bin/gmk`.

**Manual:**
```bash
go build -o gmk .
mv -v gmk ~/.local/bin/gmk
```

> **Note:** Ensure `~/.local/bin` is in your `$PATH`:
> ```bash
> echo $PATH | grep .local/bin
> ```

---

## Usage

```bash
gmk <language>
```

**Supported languages:**

| Language | Key |
|---|---|
| Go | `go` |
| Python | `python` |
| C | `c` |
| C++ | `cpp` |
| Java | `java` |
| Rust | `rust` |
| Node.js | `node` |

**Examples:**
```bash
gmk go
gmk python
gmk rust
```

Running `gmk <language>` creates the folder structure in the **current working directory**.

---

## Folder Structures

**Go**
```
cmd/ internal/ pkg/ api/ configs/ scripts/ deployments/ test/ docs/
```

**Python**
```
src/ tests/ scripts/ docs/ data/ logs/
```

**C**
```
src/ include/ lib/ tests/ build/ docs/
```

**C++**
```
src/ include/ lib/ tests/ build/ cmake/ docs/
```

**Java**
```
src/main/java/ src/main/resources/ src/test/java/ src/test/resources/ docs/ lib/
```

**Rust**
```
src/ tests/ benches/ examples/ docs/
```

**Node.js**
```
src/ tests/ public/ routes/ controllers/ models/ middleware/ configs/ docs/
```
