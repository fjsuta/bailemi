const API_BASE_URL = 'http://localhost:8080/v1';

class BailemiApp {
    constructor() {
        this.currentUser = null;
        this.accessToken = localStorage.getItem('access_token');
        this.currentSong = null;
        this.isPlaying = false;
        this.songs = [];
        this.currentSongIndex = 0;
        this.audio = new Audio();
        this.init();
    }

    init() {
        this.checkAuth();
        this.setupEventListeners();
        this.loadSongs();
        this.setupAudioPlayer();
        this.animateOnScroll();
    }

    async checkAuth() {
        if (this.accessToken) {
            try {
                const response = await fetch(`${API_BASE_URL}/user/me`, {
                    headers: {
                        'Authorization': `Bearer ${this.accessToken}`
                    }
                });
                if (response.ok) {
                    const data = await response.json();
                    this.currentUser = data.data;
                    this.updateUserUI();
                } else {
                    this.logout();
                }
            } catch (error) {
                console.error('Auth check failed:', error);
            }
        }
    }

    updateUserUI() {
        const userActions = document.querySelector('.user-actions');
        if (this.currentUser) {
            userActions.innerHTML = `
                <div class="user-info">
                    <img src="${this.currentUser.avatar || '/placeholder-avatar.png'}" alt="头像" class="user-avatar">
                    <span class="user-name">${this.currentUser.username}</span>
                </div>
                <button class="btn btn-outline" id="logoutBtn">退出</button>
            `;
            document.getElementById('logoutBtn')?.addEventListener('click', () => this.logout());
        }
    }

    setupEventListeners() {
        document.getElementById('loginBtn')?.addEventListener('click', () => this.showModal('loginModal'));
        document.getElementById('registerBtn')?.addEventListener('click', () => this.showModal('registerModal'));
        document.getElementById('toRegister')?.addEventListener('click', (e) => {
            e.preventDefault();
            this.hideModal('loginModal');
            this.showModal('registerModal');
        });
        document.getElementById('toLogin')?.addEventListener('click', (e) => {
            e.preventDefault();
            this.hideModal('registerModal');
            this.showModal('loginModal');
        });

        document.querySelectorAll('.modal-close').forEach(btn => {
            btn.addEventListener('click', () => {
                this.hideAllModals();
            });
        });

        document.getElementById('loginForm')?.addEventListener('submit', (e) => this.handleLogin(e));
        document.getElementById('registerForm')?.addEventListener('submit', (e) => this.handleRegister(e));

        document.querySelector('.verify-btn')?.addEventListener('click', (e) => this.handleSendCode(e));

        document.getElementById('playBtn')?.addEventListener('click', () => this.togglePlay());
        document.getElementById('prevBtn')?.addEventListener('click', () => this.playPrevious());
        document.getElementById('nextBtn')?.addEventListener('click', () => this.playNext());

        const searchInput = document.querySelector('.search-input');
        let searchTimeout;
        searchInput?.addEventListener('input', (e) => {
            clearTimeout(searchTimeout);
            searchTimeout = setTimeout(() => {
                this.handleSearch(e.target.value);
            }, 500);
        });
    }

    setupAudioPlayer() {
        this.audio.addEventListener('ended', () => {
            this.playNext();
        });

        this.audio.addEventListener('timeupdate', () => {
            this.updateProgress();
        });

        this.audio.addEventListener('loadedmetadata', () => {
            this.updateDuration();
        });

        const progressBar = document.querySelector('.player-progress');
        progressBar?.addEventListener('click', (e) => {
            const rect = progressBar.getBoundingClientRect();
            const percent = (e.clientX - rect.left) / rect.width;
            this.audio.currentTime = percent * this.audio.duration;
        });
    }

    updateProgress() {
        const progressFill = document.getElementById('progressFill');
        const timeCurrent = document.querySelector('.time-current');
        if (progressFill && this.audio.duration) {
            const percent = (this.audio.currentTime / this.audio.duration) * 100;
            progressFill.style.width = `${percent}%`;
        }
        if (timeCurrent) {
            timeCurrent.textContent = this.formatTime(this.audio.currentTime);
        }
    }

    updateDuration() {
        const timeTotal = document.querySelector('.time-total');
        if (timeTotal) {
            timeTotal.textContent = this.formatTime(this.audio.duration);
        }
    }

    formatTime(seconds) {
        if (isNaN(seconds)) return '00:00';
        const mins = Math.floor(seconds / 60);
        const secs = Math.floor(seconds % 60);
        return `${mins.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`;
    }

    showModal(modalId) {
        const modal = document.getElementById(modalId);
        if (modal) {
            modal.classList.add('show');
            document.body.style.overflow = 'hidden';
        }
    }

    hideModal(modalId) {
        const modal = document.getElementById(modalId);
        if (modal) {
            modal.classList.remove('show');
            document.body.style.overflow = '';
        }
    }

    hideAllModals() {
        document.querySelectorAll('.modal').forEach(modal => {
            modal.classList.remove('show');
        });
        document.body.style.overflow = '';
    }

    async handleLogin(e) {
        e.preventDefault();
        const form = e.target;
        const submitBtn = form.querySelector('button[type="submit"]');
        
        const formData = new FormData(form);
        const data = {
            login_type: formData.get('login_type'),
            account: formData.get('account'),
            password: formData.get('password')
        };

        if (!data.account || !data.password) {
            alert('请填写完整的登录信息');
            return;
        }

        submitBtn.textContent = '登录中...';
        submitBtn.disabled = true;

        try {
            const response = await fetch(`${API_BASE_URL}/auth/login`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            });

            const result = await response.json();
            if (result.code === 0) {
                this.accessToken = result.data.access_token;
                localStorage.setItem('access_token', this.accessToken);
                this.currentUser = result.data.user;
                this.updateUserUI();
                this.hideModal('loginModal');
                alert('🎉 登录成功！欢迎回来，' + this.currentUser.username);
            } else {
                alert('❌ 登录失败：' + result.message);
            }
        } catch (error) {
            console.error('Login error:', error);
            alert('❌ 登录失败，请检查网络连接');
        } finally {
            submitBtn.textContent = '登录';
            submitBtn.disabled = false;
        }
    }

    async handleRegister(e) {
        e.preventDefault();
        const form = e.target;
        const submitBtn = form.querySelector('button[type="submit"]');
        
        const formData = new FormData(form);
        const data = {
            username: formData.get('username'),
            email: formData.get('email'),
            phone: formData.get('phone'),
            password: formData.get('password'),
            verify_code: formData.get('verify_code')
        };

        if (!data.username || data.username.length < 3) {
            alert('❌ 用户名至少需要3个字符');
            return;
        }

        if (!data.email && !data.phone) {
            alert('❌ 邮箱或手机号至少填写一个');
            return;
        }

        if (data.password.length < 8) {
            alert('❌ 密码至少需要8个字符');
            return;
        }

        if (!data.verify_code) {
            alert('❌ 请输入验证码');
            return;
        }

        submitBtn.textContent = '注册中...';
        submitBtn.disabled = true;

        try {
            const response = await fetch(`${API_BASE_URL}/auth/register`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(data)
            });

            const result = await response.json();
            if (result.code === 0) {
                this.accessToken = result.data.access_token;
                localStorage.setItem('access_token', this.accessToken);
                this.currentUser = result.data.user;
                this.updateUserUI();
                this.hideModal('registerModal');
                alert('🎉 注册成功！欢迎加入百米乐，' + this.currentUser.username);
            } else {
                alert('❌ 注册失败：' + result.message);
            }
        } catch (error) {
            console.error('Register error:', error);
            alert('❌ 注册失败，请检查网络连接');
        } finally {
            submitBtn.textContent = '注册';
            submitBtn.disabled = false;
        }
    }

    async handleSendCode(e) {
        const btn = e.target;
        const registerForm = document.getElementById('registerForm');
        const email = registerForm.querySelector('input[name="email"]')?.value;
        const phone = registerForm.querySelector('input[name="phone"]')?.value;

        if (!email && !phone) {
            alert('❌ 请先填写邮箱或手机号');
            return;
        }

        btn.textContent = '发送中...';
        btn.disabled = true;

        try {
            const response = await fetch(`${API_BASE_URL}/auth/send-code`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({ email, phone })
            });

            const result = await response.json();
            if (result.code === 0) {
                alert('✅ 验证码已发送，请查收');
                this.startCountdown(btn);
            } else {
                alert('❌ 发送失败：' + result.message);
                btn.textContent = '获取验证码';
                btn.disabled = false;
            }
        } catch (error) {
            console.error('Send code error:', error);
            alert('❌ 发送失败，请稍后重试');
            btn.textContent = '获取验证码';
            btn.disabled = false;
        }
    }

    startCountdown(btn) {
        let countdown = 60;
        const originalText = '获取验证码';
        
        const timer = setInterval(() => {
            countdown--;
            if (countdown > 0) {
                btn.textContent = `${countdown}秒后重试`;
            } else {
                clearInterval(timer);
                btn.textContent = originalText;
                btn.disabled = false;
            }
        }, 1000);
    }

    logout() {
        this.accessToken = null;
        this.currentUser = null;
        localStorage.removeItem('access_token');
        this.audio.pause();
        location.reload();
    }

    async loadSongs() {
        try {
            const response = await fetch(`${API_BASE_URL}/song/hot?limit=8`);
            const result = await response.json();
            if (result.code === 0) {
                this.songs = result.data.list || result.data;
                this.renderSongs(this.songs);
                this.renderRankList(this.songs);
            }
        } catch (error) {
            console.error('Load songs failed:', error);
        }
    }

    renderSongs(songs) {
        const songList = document.getElementById('songList');
        if (!songList) return;

        songList.innerHTML = (songs.list || songs).map((song, index) => `
            <div class="song-card" data-song-id="${song.id}" data-index="${index}">
                <div class="song-cover">
                    <img src="${song.cover_url || 'https://picsum.photos/200/200?random=' + (index + 10)}" alt="${song.title}">
                    <div class="song-overlay">
                        <button class="play-btn" data-song-id="${song.id}" data-index="${index}">▶</button>
                    </div>
                </div>
                <div class="song-info">
                    <h3 class="song-title">${song.title}</h3>
                    <p class="song-artist">${typeof song.artist === 'string' ? song.artist : (song.artist?.name || '未知艺术家')}</p>
                </div>
            </div>
        `).join('');

        songList.querySelectorAll('.play-btn').forEach(btn => {
            btn.addEventListener('click', (e) => {
                e.stopPropagation();
                const index = parseInt(btn.dataset.index);
                this.playSongByIndex(index);
            });
        });

        songList.querySelectorAll('.song-card').forEach(card => {
            card.addEventListener('dblclick', () => {
                const index = parseInt(card.dataset.index);
                this.playSongByIndex(index);
            });
        });
    }

    renderRankList(songs) {
        const rankList = document.querySelector('.rank-list');
        if (!rankList) return;

        const displaySongs = (songs.list || songs).slice(0, 8);
        rankList.innerHTML = displaySongs.map((song, index) => `
            <div class="rank-item" data-index="${index}">
                <span class="rank-number ${index < 3 ? 'rank-top' : ''}">${index + 1}</span>
                <div class="rank-info">
                    <h4 class="rank-song">${song.title}</h4>
                    <p class="rank-artist">${typeof song.artist === 'string' ? song.artist : (song.artist?.name || '未知艺术家')} - ${song.album || '未知专辑'}</p>
                </div>
                <span class="rank-count">${this.formatCount(song.play_count)}</span>
            </div>
        `).join('');

        rankList.querySelectorAll('.rank-item').forEach(item => {
            item.addEventListener('dblclick', () => {
                const index = parseInt(item.dataset.index);
                this.playSongByIndex(index);
            });
        });
    }

    formatCount(count) {
        if (count >= 10000) {
            return (count / 10000).toFixed(1) + '万';
        }
        return count.toString();
    }

    async playSongByIndex(index) {
        if (index < 0 || index >= this.songs.length) return;

        this.currentSongIndex = index;
        const song = this.songs[index];

        try {
            let playUrl = song.play_url;
            
            if (!playUrl) {
                const response = await fetch(`${API_BASE_URL}/song/${song.id}/play-url`);
                const result = await response.json();
                if (result.code === 0) {
                    playUrl = result.data.play_url;
                }
            }

            this.audio.src = playUrl;
            this.audio.play();
            this.isPlaying = true;

            this.playSong({
                id: song.id,
                title: song.title,
                artist: typeof song.artist === 'string' ? song.artist : (song.artist?.name || '未知艺术家'),
                cover: song.cover_url || `https://picsum.photos/60/60?random=${song.id}`
            });

            this.updatePlayerUI();

            if (this.accessToken) {
                this.reportPlay(song.id);
            }
        } catch (error) {
            console.error('Play song failed:', error);
            alert('❌ 播放失败，请稍后重试');
        }
    }

    playSong(song) {
        this.currentSong = song;
        this.isPlaying = true;
        this.updatePlayerUI();
    }

    updatePlayerUI() {
        const playerTitle = document.querySelector('.player-title');
        const playerArtist = document.querySelector('.player-artist');
        const playerCover = document.querySelector('.player-cover');
        const playBtn = document.getElementById('playBtn');

        if (this.currentSong) {
            playerTitle.textContent = this.currentSong.title;
            playerArtist.textContent = this.currentSong.artist;
            if (this.currentSong.cover) {
                playerCover.src = this.currentSong.cover;
            }
        }

        if (playBtn) {
            playBtn.textContent = this.isPlaying ? '⏸' : '▶';
        }
    }

    togglePlay() {
        if (!this.currentSong && this.songs.length > 0) {
            this.playSongByIndex(0);
            return;
        }

        if (this.isPlaying) {
            this.audio.pause();
            this.isPlaying = false;
        } else {
            this.audio.play();
            this.isPlaying = true;
        }
        this.updatePlayerUI();
    }

    playPrevious() {
        let newIndex = this.currentSongIndex - 1;
        if (newIndex < 0) {
            newIndex = this.songs.length - 1;
        }
        this.playSongByIndex(newIndex);
    }

    playNext() {
        let newIndex = this.currentSongIndex + 1;
        if (newIndex >= this.songs.length) {
            newIndex = 0;
        }
        this.playSongByIndex(newIndex);
    }

    async reportPlay(songId) {
        try {
            await fetch(`${API_BASE_URL}/play/report`, {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${this.accessToken}`
                },
                body: JSON.stringify({
                    song_id: songId,
                    play_duration: 30,
                    total_duration: 180,
                    quality: 'high'
                })
            });
        } catch (error) {
            console.error('Report play failed:', error);
        }
    }

    async handleSearch(keyword) {
        if (!keyword.trim()) return;

        try {
            const response = await fetch(`${API_BASE_URL}/search?q=${encodeURIComponent(keyword)}`);
            const result = await response.json();
            if (result.code === 0) {
                if (result.data.songs?.list?.length > 0) {
                    this.songs = result.data.songs.list;
                    this.renderSongs(this.songs);
                }
            }
        } catch (error) {
            console.error('Search failed:', error);
        }
    }

    animateOnScroll() {
        const observer = new IntersectionObserver((entries) => {
            entries.forEach(entry => {
                if (entry.isIntersecting) {
                    entry.target.classList.add('visible');
                }
            });
        }, { threshold: 0.1 });

        document.querySelectorAll('.song-card, .playlist-card, .rank-item').forEach(el => {
            observer.observe(el);
        });
    }
}

document.addEventListener('DOMContentLoaded', () => {
    window.app = new BailemiApp();
});
