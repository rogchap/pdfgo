build-darwin_arm64:
	cd external/skia; \
	rm -rf out/mac-apple; \
	bin/gn gen out/mac-apple --args='\
	target_cpu="arm64" \
	is_official_build=true \
	skia_use_harfbuzz=true \
	skia_use_sfntly=true \
	skia_use_icu=false \
	skia_use_system_harfbuzz=false \
	skia_use_system_libjpeg_turbo=false \
	skia_use_system_libpng=false \
	skia_use_system_libwebp=false \
	skia_use_system_expat=false \
	skia_use_system_zlib=false \
	'; \
	../depot_tools/ninja -C out/mac-apple
