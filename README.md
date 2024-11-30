<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body class="bg-gray-100 text-gray-800 font-sans leading-relaxed">

  <header class="bg-gray-800 text-white py-6 text-center">
    <h1 class="text-3xl font-bold">GitHub Stats View</h1>
    <p class="mt-2 text-lg">A Go API to display statistics of languages used in your GitHub repositories.</p>
  </header>

  <main class="max-w-4xl mx-auto my-8 p-6 bg-white shadow-lg rounded-lg">
    <section class="mb-6">
      <h2 class="text-2xl font-semibold mb-4">✨ Features</h2>
      <ul class="list-disc list-inside space-y-2">
        <li>Displays the percentage of languages used in your GitHub repositories.</li>
        <li>Compatible with both public and private repositories.</li>
        <li>Easy to configure and use.</li>
      </ul>
    </section>
    <section class="mb-6">
      <h2 class="text-2xl font-semibold mb-4">⚙️ Requirements</h2>
      <p>Before starting, make sure you have:</p>
      <ul class="list-disc list-inside space-y-2">
        <li><strong>Go 1.13</strong> or a later version.</li>
        <li>A <strong>GitHub personal token</strong> (with permissions for private repositories if necessary).</li>
        <li><strong>Docker</strong> for hosting the API (or another method).</li>
      </ul>
    </section>
    <section class="mb-6">
      <h2 class="text-2xl font-semibold mb-4">🚀 Installation</h2>
      <ol class="list-decimal list-inside space-y-2">
        <li>Clone this repository:
          <pre class="bg-gray-100 p-2 rounded mt-1"><code>git clone https://github.com/lulu38m/github-language-stats.git</code></pre>
        </li>
        <li>Go to the directory:
          <pre class="bg-gray-100 p-2 rounded mt-1"><code>cd github-language-stats</code></pre>
        </li>
        <li>Set up the GitHub token as an environment variable:
          <pre class="bg-gray-100 p-2 rounded mt-1"><code>export GITHUB_TOKEN=your_github_token</code></pre>
        </li>
        <li>Run the API:
          <pre class="bg-gray-100 p-2 rounded mt-1"><code>go run main.go</code></pre>
        </li>
      </ol>
    </section>
 <!-- Docker -->
    <section class="mb-6">
  <h2 class="text-2xl font-semibold mb-4">🐳 Hosting with Docker (or another method)</h2>
  <p>You can host the API using Docker or any other method of your choice.</p>
  <ol class="list-decimal list-inside space-y-2">
    <li>Build the Docker image:
      <pre class="bg-gray-100 p-2 rounded mt-1"><code>docker build -t github-language-stats .</code></pre>
    </li>
    <li>Run the Docker container:
      <pre class="bg-gray-100 p-2 rounded mt-1"><code>docker run -e GITHUB_TOKEN=your_github_token -p 8080:8080 github-language-stats</code></pre>
    </li>
  </ol>
  <p class="mt-2">The API will be available at <code class="bg-gray-100 p-1 rounded">http://"your machine's IP":8080/stats/language</code>.</p>
  <p class="mt-2">If you prefer not to use Docker, you can also run the API directly on your machine by following the installation instructions above.</p>
</section>
    <!-- Example -->
    <section class="mb-6">
      <h2 class="text-2xl font-semibold mb-4">📄 Example Response</h2>
      <pre class="bg-gray-100 p-4 rounded">
  <img src="http://158.178.197.230:8080/stats/language" alt="GitHub Language Stats"/>
      </pre>
    </section>
    <p class="mt-4 text-gray-600">
  You can also customize the appearance of the generated SVG by modifying its style (color, size, etc.) as per your needs!
</p>
    <!-- Support -->
    <section>
      <h2 class="text-2xl font-semibold mb-4">📞 Support</h2>
      <p>If you have any questions or issues, please open an issue on the GitHub repository.</p>
    </section>
  </main>

  <!-- Footer -->
  <footer class="text-center text-sm text-gray-600 mt-6">
  <p class="mb-2">Developed by <strong>Luluberlu!</strong></p>
</footer>

</body>
</html>
