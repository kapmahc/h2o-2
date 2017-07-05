class HomeController < ApplicationController

  def index
    @cards = Card.where(loc: 'home').order(sort_order: :desc)
  end

  def dashboard
    authorize :dashboard, :sign_in?
    @cards = [
        {
            label: '.self',
            links: [
                {label: 'shared.personal-bar.profile', href: edit_user_registration_path},
                {label: 'devise.invitations.new.header', href: new_user_invitation_path},
            ]
        },
    ]
    forum_links = [
        {label: 'forum.articles.new.title', href: new_forum_article_path},
        {label: 'forum.articles.index.title', href: forum_articles_path},
        {label: 'forum.comments.index.title', href: forum_comments_path},
    ]

    if current_user.is_admin?
      @cards << {
          label: '.site',
          links: [
              {label: 'admin.site.status.title', href: admin_site_status_path},
              {label: 'admin.site.info.title', href: admin_site_info_path},
              {label: 'admin.site.nav.header', href: admin_site_nav_path(loc: :header, type: :dropdown)},
              {label: 'admin.site.nav.footer', href: admin_site_nav_path(loc: :footer, type: :link)},
              {label: 'admin.cards.index.title', href: admin_cards_path},
              {label: 'admin.site.languages.title', href: admin_site_languages_path},
              {label: 'admin.site.donate.title', href: admin_site_donate_path},

              {label: 'admin.users.index.title', href: admin_users_path},
              {label: 'leave_words.index.title', href: leave_words_path},
          ]
      }
      forum_links << {label: 'forum.tags.index.title', href: forum_tags_path}
    end

    @cards << {label: '.forum', links: forum_links}

  end

  def robots
  end

  def donate
  end

  def search

  end
end