<h1> Go / Fiber / WebSockets / Redis Streams </h1>

<hr>

<h2> Конфигурация и окружение</h2>

<p>Шаблон файла окружения лежит в <code>.env.example</code>.  
Его необходимо <b>переименовать</b> или <b>скопировать</b> в <code>.env</code>:</p>

<pre><code>cp .env.example .env</code></pre>

<p><b>Параметры по умолчанию:</b></p>
<ul>
    <li>Порт: <code>8080</code></li>
    <li>JWT secret: <code>"a-string-secret-at-least-256-bits-long"</code></li>
</ul>

<hr>

<h2> Docker</h2>

<p><b>Дефолтный порт сервера: 8080</b></p>

<p>Запуск проекта одной командой:</p>

<pre><code>docker compose up --build
</code></pre>

<p>Альтернативный вариант:</p>

<pre><code>
docker compose build app
docker compose up
</code></pre>

<p>Создать контейнер и попасть внутрь:</p>

<pre><code>docker compose run --rm app bash
</code></pre>

<hr>

<h2> Go</h2>

<p>Точка входа:</p>

<pre><code>./cmd/app/main.go
</code></pre>

<p>Собрать бинарник:</p>

<pre><code>go build -o server ./cmd/app/main.go
</code></pre>

<hr>

<h2>🧪 Тестирование</h2>

<p>В папке <code>/web</code> находится простой HTML+JS для ручной проверки</p>

<h3>Готовые JWT токены</h3>

<div class="token-box">
    <p><b>user_id = 101</b></p>
    <pre><code>eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTAxIn0.2nbyHB2XmSbtk_UfFfcP3rXjOolCSdwGkO9rtuUEexg</code></pre>
</div>

<div class="token-box">
    <p><b>user_id = 102</b></p>
    <pre><code>eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTAyIn0.IijsmEuksoCXxpbflf_Kz4zgKJ3K2tNHb9qsIHZd210</code></pre>
</div>

<h3>Особенности поведения</h3>
<p>Если нажать <b>Connect</b>, когда соединение уже открыто:</p>
<ul>
    <li>появится ошибка: «нельзя под одним токеном два коннекта»;</li>
    <li>старое соединение потеряется на клиенте;</li>
    <li>но фактически останется живым (видно в логах).</li>
</ul>

<hr>

</body>
</html>
