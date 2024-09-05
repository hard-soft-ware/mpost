## Minimum permissible structure
 - _run
    - values
        - name.txt
        - ver.txt
    - scripts
        - sys.sh

## Templates for hooks
- [commit](commit-hook.sh)
- [push](push-hook.sh)

## Constant generator templates
- [C](scripts/creator_const_C.sh)
- [G0](scripts/creator_const_Go.sh)
- [json](scripts/creator_const_json.sh)


---

# [sys.sh](scripts/sys.sh)

##### Help
```bash
scripts/sys.sh --help
```

## Reading

##### Name
```bash
scripts/sys.sh --name
```

##### Full version
```bash
scripts/sys.sh --ver
```

### Reading subversions

##### Major version
```bash
scripts/sys.sh --major
```

##### Minor version
```bash
scripts/sys.sh --minor
```

##### Patch version
```bash
scripts/sys.sh --patch
```

## Increment

##### Minor version
```bash
scripts/sys.sh --increment --minor
```

##### Patch versions
```bash
scripts/sys.sh --increment --patch
```

---

# [git.sh](scripts/git.sh)

##### Help
```bash
scripts/git.sh --help
```

## Reading

##### Name of the current branch
```bash
scripts/git.sh --branch
```

##### Name of the last tag (there will be an error if there are no tags)
```bash
scripts/git.sh --tag
```

##### Hash sum of the last commit of the current branch
```bash
scripts/git.sh --hash
```

## Working in hooks

### Adding

##### _commit-msg_ (anchoring to [commit-hook.sh](commit-hook.sh))
```bash
scripts/git.sh --add_commit
```

##### _pre-push_ (anchoring to [push-hook.sh](push-hook.sh))
```bash
scripts/git.sh --add_push
```

### Removing

##### _commit-msg_
```bash
scripts/git.sh --del_commit
```

##### _pre-push_
```bash
scripts/git.sh --del_push
```