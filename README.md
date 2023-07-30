# Geiriadur

- [Installation](#installation)
- [Mutations](#mutations)
  - [List letters](#list-letters)

## Installation

```bash
go install github.com/auenc/geiriadur
```

## Mutations

The mutations command contains subcommands that can help with mutating Welsh words:

### list-letters

```
geiriadur mutations list-letters {mutation}
```

This command will list the letters that get changed by a specified mutation.

Geiriadur supports listing letters for the following mutations:

- soft
- nasal
- aspiriate
- all

#### Example

Below is the output when listing the letters changed by the `soft` mutation

```
geiriadur mutation list-letters soft
+-------------+------+
| NO MUTATION | SOFT |
+-------------+------+
| b           | f    |
| c           | g    |
| d           | dd   |
| g           |      |
| ll          | l    |
| m           | f    |
| p           | b    |
| rh          | r    |
| t           | d    |
+-------------+------+
```

### Mutate

```
geiriadur mutations mutate {word}
```

The mutate command will display a table showing how a word would change when mutated

**Note** - The mutate command will mutate words based off of the starting letters, and doesn't care about any usuage rules. For example, a person's name will still show mutations if it starts with a relevant letter.

#### Example

Below is the output when listing the mutations for the word `cyfrifiadur`

```
geiriadur mutations mutate cyfrifiadur
+-------------+---------------+
|  MUTATION   |     WORD      |
+-------------+---------------+
| -           | cyfrifiadur   |
| Soft        | gyfrifiadur   |
| Nasal       | nghyfrifiadur |
| Aspirate    | chyfrifiadur  |
| H-Prothesis |               |
+-------------+---------------+
```
