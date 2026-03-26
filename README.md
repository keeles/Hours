# Hours

CLI tool to quickly track hours worked on different projects throughout the day/week/month/year

I always forget to note hours worked and then my PM or Lead asks how many hours a task took and I have to try to recall my day from 2 weeks ago or guesstimate a number. Not anymore! Now when I commit my code I can also run a quick terminal command to keep track of how many hours something took.

```bash
hours new client-name
hours task client-name task-name
hours add client-name task-name 240 # minutes - use -H flag to add hours
hours get client-name 
task-name: 4 hours
```

And just like that I can check to see that updating the README took 4 hours to complete.

## CLI

### New

Add a client or category to the database

- Each task belongs to a client
- Deleting a client results in all child tasks being deleted

```bash
hours new openai
```

### Delete

Delete a client or category from the database

```bash
hours delete openai
```

### Task

Add a new task under a client or category

```bash
hours task openai finish-building-AGI
```

### Add

Add time to an exisiting task

- Amount parameter defaults to minutes
- Use flag -H to insert integer value as hours

```bash
hours add openai finish-building-AGI 240 # 4 hours
hours add openai finish-building-AGI 4 -H # 4 hours
```

### Remove

Remove time from an exisiting task

- Amount parameter defaults to minutes
- Use flag -H to insert integer value as hours
- Alias: rm

```bash
hours remove openai finish-building-AGI 60 # 1 hour
hours remove openai finish-building-AGI 4 -H # 1 hour 
```

### Complete

Mark task completed for a client or category

#### DELETES TASK FROM DATABASE

```bash
hours complete openai finish-building-AGI
```

### List

List all clients/categories and their tasks

- Alias: ls

```bash
hours ls
```

### Get

List tasks for specific client or category

```bash
hours get openai
```

### Start

Starts a timer

- Timer is linked to a client on start, optionally link task
- Clients can be linked to directories on your machine allowing timers to be started without any parameters passed
- Only one timer can be running at a time

```bash
hours start openai
hours start openai finish-building-AGI

# Optionally if directory is linked with project
# eg. if ran inside of ~/Projects/openai
hours start
```

### Stop

Stops current running timer

- If no task was specified, you will be prompted to select an existing or create a new one

```bash
hours stop
```

### Time

Lists time elapsed in current running timer

```bash
hours time
```

### Config

Configuration options

#### Add Directory

Associate current directory with a client

- Alias: add-dir

```bash
# Run inside project directory
hours config add-directory openai
```

#### Remove Directory

Remove current directory association

- Alias: rm-dir

```bash
# Run inside project directory
hours config remove-directory openai
```

#### List Directories

List linked directories

- Alias: ls

```bash
hours config list
```
