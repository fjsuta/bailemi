const express = require('express');
const cors = require('cors');

const app = express();
const PORT = 8080;

app.use(cors());
app.use(express.json());

const demoUsers = [
    { id: 1, username: 'demo', nickname: '演示用户', avatar: 'https://picsum.photos/200/200?random=1', email: 'demo@example.com' }
];

const demoArtists = [
    { id: 1, name: '周杰伦', avatar_url: 'https://picsum.photos/200/200?random=10', fans_count: 10000000 },
    { id: 2, name: '林俊杰', avatar_url: 'https://picsum.photos/200/200?random=11', fans_count: 8000000 },
    { id: 3, name: '邓紫棋', avatar_url: 'https://picsum.photos/200/200?random=12', fans_count: 6000000 },
    { id: 4, name: '陈奕迅', avatar_url: 'https://picsum.photos/200/200?random=13', fans_count: 9000000 }
];

const demoAlbums = [
    { id: 1, title: '范特西', artist: '周杰伦', artist_id: 1, cover_url: 'https://picsum.photos/300/300?random=20', play_count: 5000000 },
    { id: 2, title: '曹操', artist: '林俊杰', artist_id: 2, cover_url: 'https://picsum.photos/300/300?random=21', play_count: 3000000 }
];

const demoSongs = [
    { id: 1, title: '双截棍', artist: '周杰伦', artist_id: 1, album: '范特西', album_id: 1, cover_url: 'https://picsum.photos/300/300?random=20', duration: 185, play_url: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3', play_count: 1000000, like_count: 50000 },
    { id: 2, title: '爱在西元前', artist: '周杰伦', artist_id: 1, album: '范特西', album_id: 1, cover_url: 'https://picsum.photos/300/300?random=20', duration: 245, play_url: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-2.mp3', play_count: 800000, like_count: 40000 },
    { id: 3, title: '简单爱', artist: '周杰伦', artist_id: 1, album: '范特西', album_id: 1, cover_url: 'https://picsum.photos/300/300?random=20', duration: 270, play_url: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-3.mp3', play_count: 1200000, like_count: 60000 },
    { id: 4, title: '曹操', artist: '林俊杰', artist_id: 2, album: '曹操', album_id: 2, cover_url: 'https://picsum.photos/300/300?random=21', duration: 260, play_url: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-4.mp3', play_count: 700000, like_count: 35000 },
    { id: 5, title: '江南', artist: '林俊杰', artist_id: 2, album: '曹操', album_id: 2, cover_url: 'https://picsum.photos/300/300?random=21', duration: 240, play_url: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-5.mp3', play_count: 900000, like_count: 45000 },
    { id: 6, title: '光年之外', artist: '邓紫棋', artist_id: 3, album: '光年之外', album_id: 3, cover_url: 'https://picsum.photos/300/300?random=22', duration: 235, play_url: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-6.mp3', play_count: 600000, like_count: 30000 },
    { id: 7, title: '十年', artist: '陈奕迅', artist_id: 4, album: '黑白灰', album_id: 4, cover_url: 'https://picsum.photos/300/300?random=23', duration: 205, play_url: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-7.mp3', play_count: 1500000, like_count: 75000 },
    { id: 8, title: '浮夸', artist: '陈奕迅', artist_id: 4, album: '认了吧', album_id: 5, cover_url: 'https://picsum.photos/300/300?random=24', duration: 280, play_url: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-8.mp3', play_count: 856000, like_count: 42000 }
];

const demoPlaylists = [
    { id: 1, title: '华语流行精选', description: '精选华语流行歌曲', cover_url: 'https://picsum.photos/300/300?random=30', user_id: 1, username: '演示用户', play_count: 10000, song_count: 50, is_public: true },
    { id: 2, title: '经典老歌回忆', description: '经典老歌回忆杀', cover_url: 'https://picsum.photos/300/300?random=31', user_id: 2, username: '音乐爱好者', play_count: 8000, song_count: 40, is_public: true },
    { id: 3, title: '我的私人歌单', description: '我的专属音乐', cover_url: 'https://picsum.photos/300/300?random=32', user_id: 1, username: '演示用户', play_count: 5000, song_count: 30, is_public: false }
];

app.get('/health', (req, res) => {
    res.json({ status: 'ok' });
});

app.post('/v1/auth/register', (req, res) => {
    res.json({
        code: 0, message: '注册成功', data: {
            access_token: 'demo-token', refresh_token: 'demo-refresh-token', user: demoUsers[0]
        }
    });
});

app.post('/v1/auth/login', (req, res) => {
    res.json({
        code: 0, message: '登录成功', data: {
            access_token: 'demo-token', refresh_token: 'demo-refresh-token', user: demoUsers[0]
        }
    });
});

app.get('/v1/user/me', (req, res) => {
    res.json({ code: 0, message: 'success', data: demoUsers[0] });
});

app.get('/v1/song/hot', (req, res) => {
    res.json({ code: 0, message: 'success', data: { list: demoSongs, total: demoSongs.length } });
});

app.get('/v1/song/new', (req, res) => {
    res.json({ code: 0, message: 'success', data: { list: demoSongs, total: demoSongs.length } });
});

app.get('/v1/song/:id', (req, res) => {
    const id = parseInt(req.params.id);
    const song = demoSongs.find(s => s.id === id);
    if (song) {
        res.json({ code: 0, message: 'success', data: song });
    } else {
        res.status(404).json({ code: 404, message: '歌曲不存在' });
    }
});

app.get('/v1/song/:id/play-url', (req, res) => {
    res.json({ code: 0, message: 'success', data: { play_url: 'https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3' } });
});

app.get('/v1/album/:id', (req, res) => {
    const id = parseInt(req.params.id);
    const album = demoAlbums.find(a => a.id === id);
    if (album) {
        const songs = demoSongs.filter(s => s.album_id === album.id);
        res.json({ code: 0, message: 'success', data: { ...album, songs } });
    } else {
        res.status(404).json({ code: 404, message: '专辑不存在' });
    }
});

app.get('/v1/artist/:id', (req, res) => {
    const id = parseInt(req.params.id);
    const artist = demoArtists.find(a => a.id === id);
    if (artist) {
        res.json({ code: 0, message: 'success', data: artist });
    } else {
        res.status(404).json({ code: 404, message: '歌手不存在' });
    }
});

app.get('/v1/genres', (req, res) => {
    const genres = [
        { id: 1, name: '流行', cover_url: 'https://picsum.photos/200/200?random=g1' },
        { id: 2, name: '摇滚', cover_url: 'https://picsum.photos/200/200?random=g2' },
        { id: 3, name: '民谣', cover_url: 'https://picsum.photos/200/200?random=g3' },
        { id: 4, name: '电子', cover_url: 'https://picsum.photos/200/200?random=g4' },
        { id: 5, name: '古典', cover_url: 'https://picsum.photos/200/200?random=g5' },
        { id: 6, name: '爵士', cover_url: 'https://picsum.photos/200/200?random=g6' }
    ];
    res.json({ code: 0, message: 'success', data: { list: genres, total: genres.length } });
});

app.get('/v1/playlist/recommended', (req, res) => {
    res.json({ code: 0, message: 'success', data: { list: demoPlaylists, total: demoPlaylists.length } });
});

app.get('/v1/playlist/:id', (req, res) => {
    const id = parseInt(req.params.id);
    const playlist = demoPlaylists.find(p => p.id === id);
    if (playlist) {
        res.json({ code: 0, message: 'success', data: { ...playlist, songs: demoSongs } });
    } else {
        res.status(404).json({ code: 404, message: '歌单不存在' });
    }
});

app.get('/v1/user/:id', (req, res) => {
    res.json({ code: 0, message: 'success', data: demoUsers[0] });
});

app.get('/v1/user/me/playlists', (req, res) => {
    res.json({ code: 0, message: 'success', data: { list: demoPlaylists, total: demoPlaylists.length } });
});

app.get('/v1/search', (req, res) => {
    const keyword = req.query.q || '';
    const filteredSongs = demoSongs.filter(song => song.title.includes(keyword) || song.artist.includes(keyword));
    res.json({
        code: 0,
        message: 'success',
        data: {
            songs: { list: filteredSongs, total: filteredSongs.length },
            artists: { list: demoArtists, total: demoArtists.length },
            albums: { list: demoAlbums, total: demoAlbums.length }
        }
    });
});

app.get('/v1/search/suggest', (req, res) => {
    res.json({ code: 0, message: 'success', data: ['周杰伦', '林俊杰', '邓紫棋', '陈奕迅', '流行音乐'] });
});

app.get('/v1/search/hot', (req, res) => {
    const keywords = [
        { keyword: '周杰伦', score: 100000 },
        { keyword: '林俊杰', score: 80000 },
        { keyword: '邓紫棋', score: 60000 },
        { keyword: '陈奕迅', score: 90000 },
        { keyword: '流行音乐', score: 70000 }
    ];
    res.json({ code: 0, message: 'success', data: keywords });
});

app.get('/v1/rank/:type', (req, res) => {
    res.json({ code: 0, message: 'success', data: { type: req.params.type, list: demoSongs, total: demoSongs.length } });
});

app.get('/v1/comment', (req, res) => {
    const comments = [
        { id: 1, content: '这首歌太好听了！', user_id: 1, username: '演示用户', avatar_url: 'https://picsum.photos/200/200?random=1', created_at: '2024-01-15 10:30:00', like_count: 120 },
        { id: 2, content: '百听不厌的经典', user_id: 2, username: '音乐爱好者', avatar_url: 'https://picsum.photos/200/200?random=2', created_at: '2024-01-14 15:20:00', like_count: 85 }
    ];
    res.json({ code: 0, message: 'success', data: { list: comments, total: comments.length } });
});

app.post('/v1/play/report', (req, res) => {
    res.json({ code: 0, message: '上报成功' });
});

app.get('/v1/play/history', (req, res) => {
    res.json({ code: 0, message: 'success', data: { list: demoSongs.slice(0, 3), total: 3 } });
});

app.post('/v1/comment', (req, res) => {
    res.json({ code: 0, message: '发表成功', data: { id: Date.now(), ...req.body, created_at: new Date().toISOString() } });
});

app.delete('/v1/comment/:id', (req, res) => {
    res.json({ code: 0, message: '删除成功' });
});

app.post('/v1/comment/:id/like', (req, res) => {
    res.json({ code: 0, message: '点赞成功' });
});

app.delete('/v1/comment/:id/like', (req, res) => {
    res.json({ code: 0, message: '取消点赞成功' });
});

app.post('/v1/playlist', (req, res) => {
    res.json({ code: 0, message: '创建成功', data: { id: Date.now(), ...req.body } });
});

app.put('/v1/playlist/:id', (req, res) => {
    res.json({ code: 0, message: '更新成功' });
});

app.delete('/v1/playlist/:id', (req, res) => {
    res.json({ code: 0, message: '删除成功' });
});

app.post('/v1/playlist/:id/songs', (req, res) => {
    res.json({ code: 0, message: '添加成功' });
});

app.delete('/v1/playlist/:id/songs', (req, res) => {
    res.json({ code: 0, message: '移除成功' });
});

app.put('/v1/playlist/:id/songs/sort', (req, res) => {
    res.json({ code: 0, message: '排序成功' });
});

app.post('/v1/playlist/:id/favorite', (req, res) => {
    res.json({ code: 0, message: '收藏成功' });
});

app.delete('/v1/playlist/:id/favorite', (req, res) => {
    res.json({ code: 0, message: '取消收藏成功' });
});

app.get('/v1/user/:id/following', (req, res) => {
    res.json({ code: 0, message: 'success', data: { list: [], total: 0 } });
});

app.get('/v1/user/:id/followers', (req, res) => {
    res.json({ code: 0, message: 'success', data: { list: [], total: 0 } });
});

app.post('/v1/user/:id/follow', (req, res) => {
    res.json({ code: 0, message: '关注成功' });
});

app.delete('/v1/user/:id/follow', (req, res) => {
    res.json({ code: 0, message: '取消关注成功' });
});

app.put('/v1/user/profile', (req, res) => {
    res.json({ code: 0, message: '更新成功', data: demoUsers[0] });
});

app.post('/v1/user/avatar', (req, res) => {
    res.json({ code: 0, message: '上传成功', data: { avatar_url: 'https://picsum.photos/200/200?random=' + Date.now() } });
});

app.post('/v1/auth/refresh', (req, res) => {
    res.json({ code: 0, message: '刷新成功', data: { access_token: 'demo-token', refresh_token: 'demo-refresh-token' } });
});

app.post('/v1/auth/logout', (req, res) => {
    res.json({ code: 0, message: '退出成功' });
});

app.post('/v1/auth/send-code', (req, res) => {
    res.json({ code: 0, message: '验证码已发送' });
});

app.listen(PORT, () => {
    console.log(`Demo API server running on http://localhost:${PORT}`);
});
