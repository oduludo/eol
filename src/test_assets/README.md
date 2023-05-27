# Testing assets
This folder contains assets for testing. Testing against the live API is undesirable, so unit tests will data from 
these files instead.

## Files and their purposes
**ruby_3.2.json**
Represents requesting `ruby@3.2`. The eol value has been updated to 2046, so it will always result in not having 
reached EOL.

**ruby_2.7.json**
Represents requesting `ruby@2.7`. The eol value lies in the past, so it has reached EOL.

**ruby.json**
Represents requesting all versions for `ruby`. The eol value for version 3.2 has once again been altered to 2046.
