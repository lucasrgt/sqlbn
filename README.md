<h1>GO SQL BUNDLER - Unifies multiple sql files into one</h1>

<h3>GO >= 1.22:</h3> 
<code>go install github.com/lucasrgt/sqlbn/cmd/sqlbn@latest</code>

<h3>Configuration</h3>

Create a configuration file named <code>sqlbn.yaml</code> or <code>sqlbn.yml</code>
then configure the directories of the folder containing the queries and the output query file directory and to-generate sql file name.

<b>Example:</b> <code>sqlbn.yaml</code>

```yaml
queryDir: 'internal/infrastructure/sql/queries'
outputDir: 'internal/infrastructure/sql/queries_gen.sql'
```
