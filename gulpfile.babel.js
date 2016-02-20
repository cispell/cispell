import gulp from 'gulp';
import browserSync from 'browser-sync';
import plumber from 'gulp-plumber';
import sass from 'gulp-sass';
import autoprefixer from 'gulp-autoprefixer';
import size from 'gulp-size';
import imagemin from 'gulp-imagemin';
import cache from 'gulp-cache';
import eslint from 'gulp-eslint';
import concat from 'gulp-concat';
import uglify from 'gulp-uglify';

const reload = browserSync.reload;
const mainBowerFiles = require('main-bower-files');

// Add vendor prefixes to CSS and minify it. Move minify file to the css directory.
gulp.task('css', () => {
  return gulp.src('resources/assets/sass/app.scss')
    .pipe(plumber())
    .pipe(sass.sync({
      outputStyle: 'compressed',
      precision: 10,
      includePaths: ['components/bootstrap-sass-official/assets/stylesheets'],
    }).on('error', sass.logError))
    .pipe(autoprefixer({ browsers: ['> 1%', 'last 2 versions', 'Firefox ESR'] }))
    .pipe(gulp.dest('public/css'))
    .pipe(size())
    .pipe(reload({ stream: true })
    );
});

// Copy components fonts files to the fonts directory.
gulp.task('fonts', () => {
  return gulp.src(mainBowerFiles({ filter: '**/*.{eot,svg,ttf,woff,woff2}' }))
    .pipe(gulp.dest('public/fonts'))
    .pipe(size());
});

// Minify images and move the resulting files to the images directory.
gulp.task('images', () => {
  return gulp.src('resources/assets/img/**/*')
    .pipe(cache(imagemin({
      progressive: true,
      interlaced: true,
      // don't remove IDs from SVGs, they are often used
      // as hooks for embedding and styling
      svgoPlugins: [{ cleanupIDs: false }],
    })).on('error', (err) => {
      console.log(err);
      this.end();
    }))
    .pipe(gulp.dest('public/img'))
    .pipe(size());
});

function lint(files, options) {
  return () => {
    return gulp.src(files)
      .pipe(reload({ stream: true, once: true }))
      .pipe(eslint(options))
      .pipe(eslint.format());
  };
}

gulp.task('lint', lint(['resources/assets/js/**/*.js', '!resources/assets/js/vendor/**/*']));

// Copy components javascript files to vendor folder.
gulp.task('vendor', () => {
  return gulp.src(mainBowerFiles({ checkExistence: true, filter: ['**/*.js'] }))
    .pipe(gulp.dest('resources/assets/js/vendor'))
    .pipe(size());
});

// Concatenate together all js from vendor into one minify file for the application.
gulp.task('concat', ['vendor'], () => {
  return gulp.src([
    'resources/assets/js/vendor/jquery.js',
    'resources/assets/js/vendor/bootstrap.js',
    'resources/assets/js/vendor/**/*.js',
    'resources/assets/js/**/*.js',
  ])
    .pipe(concat('app.js'))
    .pipe(uglify()).on('error', (err) => {
      console.log(err);
      this.end();
    })
    .pipe(gulp.dest('public/js'))
    .pipe(size());
});

gulp.task('js', ['concat']);

// Watch for any changes and reload the browser.
gulp.task('watch', ['build'], () => {
  browserSync({
    options: {
      proxy: 'localhost:80',
    },
  });

  gulp.watch('resources/assets/sass/**/*.scss', ['css']);
  gulp.watch('resources/assets/js/**/*.js', ['js']);
  gulp.watch('resources/assets/img/**/**', ['images']);
  gulp.watch('bower.json', ['bower:install']);
  gulp.watch('package.json', ['npm:install']);

  gulp.watch([
    'public/css/**/*.css',
    'public/fonts/**/*',
    'public/img/**/*',
    'public/js/**/*.js',
  ]).on('change', reload);
});

// Runs all commands to build the application.
// Returns the size of all assets except for the upload folder.
gulp.task('build', ['css', 'fonts', 'images', 'js'], () => {
  return gulp.src(['public/**/*', '!public/uploads'])
    .pipe(size({ title: 'build', gzip: true }));
});

// Default command for gulp. Runs the build command.
gulp.task('default', ['build']);
