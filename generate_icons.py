#!/usr/bin/env python3
"""Generate high-quality icons for Bailemi music platform using Seedream"""

import asyncio
import sys
import os

sys.path.append("/data/user/skills/byted-seedream-image-generate/scripts")
from seedream_image_generate import seedream_generate

async def generate_icons():
    print("🎨 Generating Bailemi Music Platform Icons with Seedream 5.0-lite...")
    
    prompts = [
        # 1. 主Logo - 音乐符号+紫色渐变
        {
            "prompt": "Professional music platform logo icon, minimalist musical note with purple to blue gradient, modern design, vector style, white background, clean lines, 2D flat design, high quality",
            "size": "1024x1024",
            "watermark": False,
            "output_format": "png"
        },
        
        # 2. 播放按钮图标
        {
            "prompt": "Play button icon, minimalist, vector style, white play symbol on purple gradient background, modern UI design, 2D flat, clean, high quality",
            "size": "1024x1024",
            "watermark": False,
            "output_format": "png"
        },
        
        # 3. 暂停按钮图标
        {
            "prompt": "Pause button icon, minimalist, vector style, two vertical white bars on purple background, modern UI design, 2D flat, clean",
            "size": "1024x1024",
            "watermark": False,
            "output_format": "png"
        },
        
        # 4. 搜索图标
        {
            "prompt": "Magnifying glass search icon, minimalist vector style, white on dark background, modern UI design, 2D flat",
            "size": "1024x1024",
            "watermark": False,
            "output_format": "png"
        },
        
        # 5. 用户头像图标
        {
            "prompt": "User profile avatar icon, minimalist human figure, vector style, modern UI design, 2D flat, purple theme",
            "size": "1024x1024",
            "watermark": False,
            "output_format": "png"
        },
        
        # 6. 收藏/爱心图标
        {
            "prompt": "Heart icon, minimalist vector style, white heart on pink gradient background, modern UI design, 2D flat",
            "size": "1024x1024",
            "watermark": False,
            "output_format": "png"
        },
        
        # 7. 分享图标
        {
            "prompt": "Share icon, minimalist vector, three connected dots with lines, modern UI design, 2D flat, purple theme",
            "size": "1024x1024",
            "watermark": False,
            "output_format": "png"
        },
        
        # 8. 设置图标
        {
            "prompt": "Settings gear icon, minimalist vector style, modern UI design, 2D flat, dark background with white gear",
            "size": "1024x1024",
            "watermark": False,
            "output_format": "png"
        },
        
        # 9. 上传图标
        {
            "prompt": "Upload icon, minimalist vector, arrow pointing up with cloud, modern UI design, 2D flat, purple theme",
            "size": "1024x1024",
            "watermark": False,
            "output_format": "png"
        },
        
        # 10. 主题切换图标
        {
            "prompt": "Sun and moon icon for theme toggle, minimalist vector style, half sun half moon, modern UI design",
            "size": "1024x1024",
            "watermark": False,
            "output_format": "png"
        }
    ]
    
    results = await seedream_generate(prompts, version="5.0")
    
    print("\n✅ Icons generated successfully!")
    for i, result in enumerate(results):
        print(f"\nIcon {i+1}:")
        print(f"  URL: {result.get('url', 'N/A')}")
        print(f"  Prompt: {prompts[i]['prompt'][:50]}...")
    
    return results

if __name__ == "__main__":
    asyncio.run(generate_icons())
