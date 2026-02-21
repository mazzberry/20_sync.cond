<h1>Go Concurrency Example – sync.Cond Broadcast Demo</h1>

<h2>📌 Overview</h2>

<p>
This project demonstrates how to use <strong>sync.Cond (condition variables)</strong> in Go to coordinate multiple goroutines.
</p>

<p>
The program simulates 1000 users waiting to start streaming.
When the number of ready users reaches <strong>55</strong>, all waiting goroutines are notified simultaneously using <code>Broadcast()</code>.
</p>

<p>This example helps understand:</p>

<ul>
  <li>Condition variables in Go</li>
  <li>Goroutine synchronization</li>
  <li><code>Wait()</code> and <code>Broadcast()</code> behavior</li>
  <li>Proper mutex locking with <code>sync.Cond</code></li>
</ul>

<hr>

<h2>🧠 How It Works</h2>

<h3>1️⃣ Main Function</h3>

<ul>
  <li>Creates a <code>sync.Cond</code> with a mutex.</li>
  <li>Spawns 1000 goroutines (<code>NewRequest</code>).</li>
  <li>Waits 10 seconds to allow execution.</li>
</ul>

<pre><code class="language-go">
condition := sync.NewCond(&sync.Mutex{})
</code></pre>

<hr>

<h3>2️⃣ NewRequest()</h3>

<p>Each goroutine:</p>

<ol>
  <li>Calls <code>Checking()</code></li>
  <li>Locks the condition mutex</li>
  <li>Waits until <code>ready == true</code></li>
  <li>Starts streaming after being notified</li>
</ol>

<pre><code class="language-go">
for !ready {
    condition.Wait()
}
</code></pre>

<p><strong>Important:</strong></p>
<ul>
  <li><code>Wait()</code> automatically unlocks the mutex while waiting.</li>
  <li>When awakened, it re-locks the mutex before continuing.</li>
</ul>

<hr>

<h3>3️⃣ Checking()</h3>

<p>Each user:</p>

<ul>
  <li>Prints a waiting message</li>
  <li>Sleeps 150ms (simulating validation work)</li>
  <li>Appends itself to <code>userList</code></li>
  <li>When <code>len(userList) == 55</code>:
    <ul>
      <li>Sets <code>ready = true</code></li>
      <li>Calls <code>Broadcast()</code> to wake all waiting goroutines</li>
    </ul>
  </li>
</ul>

<pre><code class="language-go">
if len(userList) == 55 {
    ready = true
    condition.Broadcast()
}
</code></pre>

<hr>

<h2>🔄 Execution Flow</h2>

<ol>
  <li>1000 goroutines start.</li>
  <li>Each waits for the <code>ready</code> signal.</li>
  <li>When the 55th user arrives:
    <ul>
      <li><code>ready</code> becomes true.</li>
      <li><code>Broadcast()</code> wakes all waiting goroutines.</li>
    </ul>
  </li>
  <li>All users start streaming at once.</li>
</ol>

<hr>

<h2>🛠 Key Concepts</h2>

<h3>sync.Cond</h3>
<p>A condition variable used for signaling between goroutines.</p>

<h3>Wait()</h3>
<ul>
  <li>Unlocks the mutex</li>
  <li>Suspends execution</li>
  <li>Re-locks mutex when awakened</li>
</ul>

<h3>Broadcast()</h3>
<p>Wakes <strong>all</strong> goroutines waiting on the condition.</p>

<h3>Why use <code>for</code> instead of <code>if</code>?</h3>

<pre><code class="language-go">
for !ready {
    condition.Wait()
}
</code></pre>

<p>
Using <code>for</code> prevents issues from spurious wakeups and race conditions.
This is a best practice in concurrent programming.
</p>

<hr>

<h2>⚠️ Important Notes</h2>

<ul>
  <li><code>userList</code> and <code>ready</code> are shared variables.</li>
  <li>They are protected using the condition’s mutex.</li>
  <li>Always lock before modifying shared state.</li>
  <li>Always use a <code>for</code> loop with <code>Wait()</code>.</li>
</ul>

<hr>

<h2>🚀 How To Run</h2>

<pre><code class="language-bash">
go run main.go
</code></pre>

<hr>

<h2>👨‍💻 Author</h2>

<p>
Mohammadreza Hosseini<br>
Go Developer | Backend Enthusiast
</p>