<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
</head>
<body class="bg-gray-100 text-gray-800 font-sans leading-relaxed">

  <header class="bg-gray-800 text-white py-6 text-center">
    <h1 class="text-3xl font-bold">GitHub Language Stats API</h1>
    <p class="mt-2 text-lg">Une API en Go pour afficher les statistiques des langages utilisés dans vos dépôts GitHub.</p>
  </header>

  <main class="max-w-4xl mx-auto my-8 p-6 bg-white shadow-lg rounded-lg">
    <section class="mb-6">
      <h2 class="text-2xl font-semibold mb-4">✨ Fonctionnalités</h2>
      <ul class="list-disc list-inside space-y-2">
        <li>Affiche le pourcentage des langages utilisés dans vos dépôts GitHub.</li>
        <li>Compatible avec les dépôts publics et privés.</li>
        <li>Facile à configurer et à utiliser.</li>
      </ul>
    </section>
    <section class="mb-6">
      <h2 class="text-2xl font-semibold mb-4">⚙️ Prérequis</h2>
      <p>Avant de commencer, assurez-vous d'avoir :</p>
      <ul class="list-disc list-inside space-y-2">
        <li><strong>Go 1.13</strong> ou une version ultérieure.</li>
        <li>Un <strong>token personnel GitHub</strong> (avec les permissions pour les dépôts privés si nécessaire).</li>
        <li><strong>Docker</strong> pour l'hébergement de l'API (ou autre méthode).</li>
      </ul>
    </section>
    <section class="mb-6">
      <h2 class="text-2xl font-semibold mb-4">🚀 Installation</h2>
      <ol class="list-decimal list-inside space-y-2">
        <li>Clonez ce dépôt :
          <pre class="bg-gray-100 p-2 rounded mt-1"><code>git clone https://github.com/lulu38m/github-language-stats.git</code></pre>
        </li>
        <li>Accédez au dossier :
          <pre class="bg-gray-100 p-2 rounded mt-1"><code>cd github-language-stats</code></pre>
        </li>
        <li>Configurez le token GitHub dans une variable d'environnement :
          <pre class="bg-gray-100 p-2 rounded mt-1"><code>export GITHUB_TOKEN=your_github_token</code></pre>
        </li>
        <li>Démarrez l'API :
          <pre class="bg-gray-100 p-2 rounded mt-1"><code>go run main.go</code></pre>
        </li>
      </ol>
    </section>
 <!-- Docker -->
    <section class="mb-6">
  <h2 class="text-2xl font-semibold mb-4">🐳 Hébergement avec Docker (ou autre méthode)</h2>
  <p>Vous pouvez héberger l'API en utilisant Docker ou toute autre méthode de votre choix.</p>
  <ol class="list-decimal list-inside space-y-2">
    <li>Construisez l'image Docker :
      <pre class="bg-gray-100 p-2 rounded mt-1"><code>docker build -t github-language-stats .</code></pre>
    </li>
    <li>Démarrez le conteneur Docker :
      <pre class="bg-gray-100 p-2 rounded mt-1"><code>docker run -e GITHUB_TOKEN=your_github_token -p 8080:8080 github-language-stats</code></pre>
    </li>
  </ol>
  <p class="mt-2">L'API sera disponible sur <code class="bg-gray-100 p-1 rounded">http://"ip de votre machine":8080/stats/language</code>.</p>
  <p class="mt-2">Si vous préférez ne pas utiliser Docker, vous pouvez également exécuter l'API directement sur votre machine en suivant les instructions d'installation ci-dessus.</p>
</section>
    <!-- Example -->
    <section class="mb-6">
      <h2 class="text-2xl font-semibold mb-4">📄 Exemple de réponse</h2>
      <pre class="bg-gray-100 p-4 rounded">
  <img src="http://158.178.197.230:8080/stats/language" alt="GitHub Language Stats"/>
      </pre>
    </section>
    <p class="mt-4 text-gray-600">
  Vous pouvez également personnaliser l'apparence du SVG généré en modifiant son style (couleur, taille, etc.) selon vos besoins !
</p>
    <!-- Support -->
    <section>
      <h2 class="text-2xl font-semibold mb-4">📞 Support</h2>
      <p>Si vous avez des questions ou des problèmes, ouvrez une issue sur le dépôt GitHub.</p>
    </section>
  </main>

  <!-- Footer -->
  <footer class="text-center text-sm text-gray-600 mt-6">
  <p class="mb-2">Développé par <strong>Luluberlu !</strong></p>
</footer>

</body>
</html>
