package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GCwdLQIYkACIgycDMXlienpqUX6UFovqzg/LzSElYHRT5s34U1EfLa0u0D9+eUR62oLHO54ljk5+puKXFJbyOvpKx91ec37Y8t5y7fPvruXw/yF4oNtmRmPHeY4ycw6ZqK5LGD3QtarvAfj3cqD9rnuvPpS6+5brUPv5DU+TTb5ffv+uXN/0o6tnr1cwOmEUB63k0CYcO0anlM3Jk2KfPU6S7d968b3zNtnnYja8josT/Wi1dZZL57FW9sK6Kn7CEml3lhtF/vcosWf+YDNri6p3u3nrk0/N8ey+5f1bI14l5Pn7b/xzbd4eS/v192Lq+8e/re4tkXypH/gnkknGIyOcTptqXm5t3CKt3JnrOXDSVEdBSvKxJ/3vw1qmMF+s6rs47LHW27n7m64NWmTkYyTzxW2yAP3vhnPfnvpgtaGuYb3mR45qRx4c3vn9/ALO73FL06KvW58h0Hv0w6zN217LVyecnQn3hSafLJgVfvPgN5zBdbL495MazFoPNDDeuRoZNbc7Ca3MNZtUas/eb3xf5jjKHDhZoeLTNuM/JTt8jK56v1yi8uvpvFOaU4wEVLYEh464WjV3R0dJk7/HjQ9W8C0povDXzVkh7nlmvScfAu5NZM7Ut4cktyZPV/cZNec+pVsqV+uX1wXsnbytL6wIL60I52tWm8ipyod6LWcKH1sjV/nq5/9dme2rT1XwTqXL0aqc8edjWoCx4Vf1QczaH0z/tjzcv/+vdPvveN7VzZ5bT57/L8pJfO/l5tVu+9+uba50nLGmqVFGqYlbq2O74r3Pq3yrr/rve9HlLpB2+aJOo2z+efFRmkz/RB9YN7BHd4hv3bD6YLaWTfcHhfnSnVovDJ9levjXLhDTphv3lWXScc/9GvaLJFj2fZtkfq0nXpTP7MFbN2TJmfF+HVqhldw4T31Uye85r+wr5m4p+bD+i+dqW56aZW/7Ld3fLSet4Pt2d2DKaZl7F8zv5yQKk75tHLp+fCltSd4l03xuL6uUfv0eq9jmsx6nYqiYoedtmh+P9Mp9ZH7pOeT8MWXLj6x1/66x6nvdFZXrK/4glNnZruvfHCh+3c+3wTvvZrXNaZ+Wv22vbm5+d/kaG7ur2lVv+OKnv+WyjKZvPw8+37b1Qv9H1XVckqdZSr7tWXhF+7lnioHr/uKx96w1Tk7d9Gb1x/KG78c1H9p/pt77/3ySda7X95NP3mN/+Q/v/cfVj/5nPf+6d+HRf+P/v3y8Vl27Vlmm8mrF35I/2Zyd9f1+l/tH0+2+S1R/5p5ZaPb+q/lPnOzf6iYc/AYFzGeYt44Y/npjTmBfnbB/6umOB6uvMX0pve5Vfm7kOfnGRgY/v8P8GbnyPx78/ssZgaGehEGBlgmZ8DI5OyITA7P1yDdyGoCvBmZRJgRhQSyyaBCAga2NYJIvEUGwijsToEAAYb/jk+ZGbA4jJUNJM/EwMTQycDAoMwC4gECAAD//71qzVnCBAAA"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
