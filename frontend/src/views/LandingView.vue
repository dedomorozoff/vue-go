<script setup>
import { useAuthStore } from '../stores/auth'
import { Shield, Zap, Sparkles, Globe, Server, Code, ArrowRight, LayoutDashboard } from 'lucide-vue-next'

const auth = useAuthStore()
</script>

<template>
  <div class="landing">
    <!-- Navbar -->
    <nav class="navbar">
      <div class="container">
        <div class="logo">
          <Zap class="icon-brand" />
          <span>Vue-Go</span>
        </div>
        <div class="nav-links">
          <a href="#features">Features</a>
          <a href="#stack">Stack</a>
          <RouterLink to="/admin" class="btn-primary">Go to Admin</RouterLink>
        </div>
      </div>
    </nav>

    <!-- Hero Section -->
    <header class="hero">
      <div class="container">
        <div class="hero-content">
          <h1 class="gradient-text">Build Fast. Ship Faster.</h1>
          <p class="hero-subtitle">
            The ultimate starter kit combining the power of Go with the elegance of Vue 3. 
            Everything you need for your next big idea, already integrated and ready to go.
          </p>
          <div class="hero-actions">
            <RouterLink to="/admin" class="btn-primary btn-lg">
              Get Started
              <ArrowRight :size="20" />
            </RouterLink>
            <a href="https://github.com/dedomorozoff/vue-go" class="btn-secondary btn-lg">View Source</a>
          </div>
        </div>
        <div class="hero-visual">
          <div class="code-card glass-card">
            <div class="code-header">
              <div class="dot red"></div>
              <div class="dot yellow"></div>
              <div class="dot green"></div>
            </div>
            <pre><code><span class="keyword">func</span> <span class="func">main</span>() {
  <span class="comment">// Pure speed</span>
  server := go.NewServer()
  server.Run(<span class="string">":8080"</span>)
}</code></pre>
          </div>
        </div>
      </div>
    </header>

    <!-- Features Section -->
    <section id="features" class="features">
      <div class="container">
        <div class="section-header">
          <span class="badge">Features</span>
          <h2>Everything in one place</h2>
        </div>
        <div class="features-grid">
          <div class="feature-card glass-card">
            <Shield class="feature-icon" color="#6366f1" />
            <h3>Secure by default</h3>
            <p>Built-in security best practices from Go and industry standards.</p>
          </div>
          <div class="feature-card glass-card">
            <Sparkles class="feature-icon" color="#f59e0b" />
            <h3>Modern UI</h3>
            <p>Clean, responsive and premium design with Vue 3 and CSS variables.</p>
          </div>
          <div class="feature-card glass-card">
            <Zap class="feature-icon" color="#10b981" />
            <h3>High Performance</h3>
            <p>Go backend ensures sub-millisecond responses for your API.</p>
          </div>
        </div>
      </div>
    </section>

    <!-- Stack Section -->
    <section id="stack" class="stack">
      <div class="container">
        <div class="section-header">
          <span class="badge">Technology Stack</span>
          <h2>Modern & Robust Tools</h2>
        </div>
        <div class="stack-grid">
          <div class="stack-item">
            <div class="stack-box glass-card">
              <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/go/go-original.svg" alt="Go" width="60" />
            </div>
            <h4>Go Backend</h4>
            <p>Go 1.25, chi Router, GORM, SQLite</p>
          </div>
          <div class="stack-item">
            <div class="stack-box glass-card">
              <img src="https://raw.githubusercontent.com/devicons/devicon/master/icons/vuejs/vuejs-original.svg" alt="Vue" width="60" />
            </div>
            <h4>Vue Frontend</h4>
            <p>Vue 3, Vite, Pinia, Vue Router</p>
          </div>
          <div class="stack-item">
            <div class="stack-box glass-card">
              <div class="custom-stack-icon"><Code :size="40" color="#64748b" /></div>
            </div>
            <h4>Ecosystem</h4>
            <p>JWT Auth, Lucide Icons, Axios, Makefile</p>
          </div>
        </div>
      </div>
    </section>

    <footer class="footer">
      <div class="container">
        <p>&copy; 2026 Vue-Go Project.</p>
      </div>
    </footer>

    <!-- Floating Action Button -->
    <RouterLink v-if="auth.isAuthenticated" to="/admin" class="fab glass-card" title="Back to Admin">
      <LayoutDashboard :size="24" />
      <span class="fab-text">Dashboard</span>
    </RouterLink>
  </div>
</template>

<style scoped>
.landing {
  background-color: #ffffff;
  color: var(--text-main);
  overflow-x: hidden;
}

.container {
  max-width: 1200px;
  margin: 0 auto;
  padding: 0 24px;
}

/* Navbar */
.navbar {
  height: 80px;
  display: flex;
  align-items: center;
  position: fixed;
  top: 0;
  width: 100%;
  background: rgba(255, 255, 255, 0.8);
  backdrop-filter: blur(12px);
  z-index: 100;
  border-bottom: 1px solid var(--border);
}

.navbar .container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.logo {
  display: flex;
  align-items: center;
  gap: 10px;
  font-weight: 800;
  font-size: 1.5rem;
  color: var(--text-main);
}

.icon-brand {
  color: var(--primary);
}

.nav-links {
  display: flex;
  align-items: center;
  gap: 32px;
}

.nav-links a:not(.btn-primary) {
  text-decoration: none;
  color: var(--text-muted);
  font-weight: 500;
  transition: color 0.2s;
}

.nav-links a:not(.btn-primary):hover {
  color: var(--primary);
}

/* Hero Section */
.hero {
  padding: 160px 0 100px;
  background: radial-gradient(circle at top right, rgba(99, 102, 241, 0.05), transparent),
              radial-gradient(circle at bottom left, rgba(245, 158, 11, 0.05), transparent);
}

.hero .container {
  display: flex;
  align-items: center;
  gap: 60px;
}

.hero-content {
  flex: 1;
}

.gradient-text {
  font-size: 4rem;
  font-weight: 800;
  line-height: 1.1;
  margin-bottom: 24px;
  background: linear-gradient(135deg, #6366f1 0%, #a855f7 50%, #f43f5e 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
}

.hero-subtitle {
  font-size: 1.25rem;
  color: var(--text-muted);
  line-height: 1.6;
  margin-bottom: 40px;
}

.hero-actions {
  display: flex;
  gap: 16px;
}

.btn-lg {
  padding: 16px 32px;
  font-size: 1.125rem;
  display: flex;
  align-items: center;
  gap: 10px;
}

.btn-secondary {
  background: white;
  border: 1px solid var(--border);
  color: var(--text-main);
  text-decoration: none;
  border-radius: 12px;
  font-weight: 600;
  transition: all 0.2s;
}

.btn-secondary:hover {
  background: var(--bg-main);
  border-color: var(--text-muted);
}

.hero-visual {
  flex: 1;
  perspective: 1000px;
}

.code-card {
  padding: 24px;
  transform: rotate3d(1, -1, 0, 15deg);
  transition: transform 0.5s;
}

.code-card:hover {
  transform: rotate3d(0, 0, 0, 0deg);
}

.code-header {
  display: flex;
  gap: 8px;
  margin-bottom: 20px;
}

.dot {
  width: 12px;
  height: 12px;
  border-radius: 50%;
}

.red { background: #ff5f56; }
.yellow { background: #ffbd2e; }
.green { background: #27c93f; }

pre {
  background: #1e293b;
  color: #f8fafc;
  padding: 20px;
  border-radius: 12px;
  font-family: 'Fira Code', monospace;
  font-size: 0.9rem;
}

.keyword { color: #f472b6; }
.func { color: #38bdf8; }
.string { color: #4ade80; }
.comment { color: #64748b; }

/* Features */
.features {
  padding: 100px 0;
}

.section-header {
  text-align: center;
  margin-bottom: 60px;
}

.badge {
  background: #f0f1ff;
  color: var(--primary);
  padding: 6px 16px;
  border-radius: 20px;
  font-size: 0.875rem;
  font-weight: 600;
}

.section-header h2 {
  font-size: 2.5rem;
  margin-top: 16px;
}

.features-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 32px;
}

.feature-card {
  padding: 40px;
  text-align: center;
  transition: transform 0.3s;
}

.feature-card:hover {
  transform: translateY(-10px);
}

.feature-icon {
  margin-bottom: 24px;
}

.feature-card h3 {
  font-size: 1.5rem;
  margin-bottom: 16px;
}

.feature-card p {
  color: var(--text-muted);
}

/* Stack Section */
.stack {
  padding: 100px 0;
  background: #fcfcfd;
}

.stack-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 40px;
}

.stack-item {
  text-align: center;
}

.stack-box {
  width: 120px;
  height: 120px;
  margin: 0 auto 24px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 24px;
  transition: transform 0.3s;
}

.stack-item:hover .stack-box {
  transform: scale(1.1) rotate(5deg);
}

.stack-item h4 {
  font-size: 1.25rem;
  margin-bottom: 8px;
}

.stack-item p {
  color: var(--text-muted);
  font-size: 0.875rem;
}

.custom-stack-icon {
  background: #f1f5f9;
  width: 60px;
  height: 60px;
  border-radius: 12px;
  display: flex;
  align-items: center;
  justify-content: center;
}

/* Floating Action Button */
.fab {
  position: fixed;
  bottom: 40px;
  right: 40px;
  background: var(--primary);
  color: white;
  padding: 16px;
  border-radius: 50px;
  display: flex;
  align-items: center;
  gap: 12px;
  text-decoration: none;
  box-shadow: 0 12px 32px rgba(99, 102, 241, 0.4);
  z-index: 1000;
  border: 1px solid rgba(255, 255, 255, 0.2);
  transition: all 0.3s cubic-bezier(0.175, 0.885, 0.32, 1.275);
  overflow: hidden;
  max-width: 56px;
}

.fab:hover {
  max-width: 200px;
  padding: 16px 24px;
  transform: translateY(-5px);
  background: var(--primary-hover);
}

.fab-text {
  font-weight: 600;
  white-space: nowrap;
  opacity: 0;
  transition: opacity 0.2s;
}

.fab:hover .fab-text {
  opacity: 1;
}

.footer {
  padding: 40px 0;
  border-top: 1px solid var(--border);
  text-align: center;
  color: var(--text-muted);
}

@media (max-width: 900px) {
  .hero .container {
    flex-direction: column;
    text-align: center;
  }
  .hero-visual {
    width: 100%;
  }
  .features-grid {
    grid-template-columns: 1fr;
  }
  .gradient-text {
    font-size: 2.5rem;
  }
}
</style>
